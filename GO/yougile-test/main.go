package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type AuthRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type KeyRequest struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	CompanyID string `json:"companyId"`
}

type KeyResponse struct {
	Key string `json:"key"`
}

type ListResponse[T any] struct {
	Paging struct {
		Count  int  `json:"count"`
		Limit  int  `json:"limit"`
		Offset int  `json:"offset"`
		Next   bool `json:"next"`
	} `json:"paging"`
	Content []T `json:"content"`
}

type Company struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	IsAdmin bool   `json:"isAdmin"`
}

type Project struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type ProjectNode struct {
	ID     string
	Title  string
	Boards []BoardNode
}

type Board struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	ProjectID string `json:"projectId"`
}

type BoardNode struct {
	ID      string
	Title   string
	Columns []ColumnNode
}

type ColumnNode struct {
	ID    string
	Title string
	Tasks []Task
}

type ColumnDetail struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	BoardID string `json:"boardId"`

	Position *int `json:"position,omitempty"`
	Index    *int `json:"index,omitempty"`
	Order    *int `json:"order,omitempty"`
}

type Task struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	ColumnID string `json:"columnId,omitempty"`
	// доп поля
	// Description string `json:"description,omitempty"`
	// Completed bool `json:"completed,omitempty"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Файл .env не найден")
	}
}

func main() {
	domain := os.Getenv("domain")
	login := os.Getenv("login")
	password := os.Getenv("password")
	apiKey := os.Getenv("apikey")

	if domain == "" || login == "" || password == "" {
		fmt.Println(".env не найден")
		return
	}

	// 1) Компании
	companies, err := getCompanies(domain, login, password)
	if err != nil {
		log.Fatalf("Ошибка компаний: %v", err)
	}
	if len(companies) == 0 {
		log.Fatalf("Компаний нет")
	}
	company := companies[0]
	//fmt.Printf("КОМПАНИЯ: %s (ID: %s)\n", company.Name, company.ID)

	// 2) Ключ
	if apiKey == "" {
		apiKey, err = getAPIKey(domain, login, password, company.ID)
		if err != nil {
			log.Fatalf("Ошибка ключа: %v", err)
		}
		if err := upsertEnvKey(".env", "apikey", apiKey); err != nil {
			log.Printf("Не смог записать apikey в .env: %v", err)
		}
	}

	// 3) Проекты (твоя функция)
	projects, err := getProjects(domain, apiKey)
	if err != nil {
		log.Fatalf("Ошибка проектов: %v", err)
	}

	// 4) Доски
	boards, err := getBoards(domain, apiKey)
	if err != nil {
		log.Fatalf("Ошибка досок: %v", err)
	}
	//fmt.Printf("Досок: %d\n", len(boards))

	// 5) Колонки (все -> детали)
	columnIDs, err := getAllColumnIDs(domain, apiKey)
	if err != nil {
		log.Fatalf("Ошибка списка колонок: %v", err)
	}
	columnsByBoard, err := getColumnsByBoard(domain, apiKey, columnIDs)
	if err != nil {
		log.Fatalf("Ошибка деталей колонок: %v", err)
	}

	// 6) Тикеты/задачи: для каждой колонки берём tasks и складываем
	tasksByColumn := map[string][]Task{}
	for boardID, cols := range columnsByBoard {
		_ = boardID
		for _, col := range cols {
			tasks, err := getTasksByColumn(domain, apiKey, col.ID)
			if err != nil {
				// чтобы не ронять весь вывод, просто показываем предупреждение
				fmt.Printf("Предупреждение: tasks для column %s: %v\n", col.ID, err)
				continue
			}
			tasksByColumn[col.ID] = tasks
		}
	}

	// 7) Собираем и печатаем иерархию: Проект (из projectId) -> Доска -> Колонка -> Тикеты
	tree := buildHierarchy(projects, boards, columnsByBoard, tasksByColumn)
	printHierarchy(company, tree)

}

func buildHierarchy(
	projects []Project,
	boards []Board,
	columnsByBoard map[string][]ColumnDetail,
	tasksByColumn map[string][]Task,
) []ProjectNode {

	// 1) Справочник проектов по id (для названий)
	pm := map[string]*ProjectNode{}
	for _, p := range projects {
		pm[p.ID] = &ProjectNode{
			ID:    p.ID,
			Title: p.Title,
		}
	}

	// 2) Если встретятся доски с projectId, которого нет в projects — добавим “заглушку”
	ensureProject := func(projectID string) *ProjectNode {
		if projectID == "" {
			projectID = "no-project"
		}
		if pm[projectID] == nil {
			title := "Проект " + projectID
			if projectID == "no-project" {
				title = "Без проекта"
			}
			pm[projectID] = &ProjectNode{ID: projectID, Title: title}
		}
		return pm[projectID]
	}

	// 3) Разложить доски по проектам
	for _, b := range boards {
		pnode := ensureProject(b.ProjectID)

		bn := BoardNode{ID: b.ID, Title: b.Title}

		cols := columnsByBoard[b.ID]
		sortColumns(cols)

		for _, c := range cols {
			cn := ColumnNode{ID: c.ID, Title: c.Title}
			cn.Tasks = tasksByColumn[c.ID]
			bn.Columns = append(bn.Columns, cn)
		}

		pnode.Boards = append(pnode.Boards, bn)
	}

	// 4) Стабильная сортировка проектов и досок
	out := make([]ProjectNode, 0, len(pm))
	for _, p := range pm {
		sort.SliceStable(p.Boards, func(i, j int) bool {
			return strings.ToLower(p.Boards[i].Title) < strings.ToLower(p.Boards[j].Title)
		})
		out = append(out, *p)
	}

	sort.SliceStable(out, func(i, j int) bool {
		return strings.ToLower(out[i].Title) < strings.ToLower(out[j].Title)
	})

	return out
}

func printHierarchy(company Company, projects []ProjectNode) {
	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("Компания: %s (%s)\n", company.Name, company.ID)
	fmt.Printf("Проектов: %d\n", len(projects))
	fmt.Println(strings.Repeat("=", 80))

	for _, p := range projects {
		fmt.Printf("\nПРОЕКТ: %s (ID: %s)\n", p.Title, p.ID)

		if len(p.Boards) == 0 {
			fmt.Println("  Доски: (0)")
			continue
		}

		for _, b := range p.Boards {
			fmt.Printf("  ДОСКА: %s (ID: %s)\n", b.Title, b.ID)

			if len(b.Columns) == 0 {
				fmt.Println("    Колонки: (0)")
				continue
			}

			for _, c := range b.Columns {
				fmt.Printf("    КОЛОНКА: %s (ID: %s)\n", c.Title, c.ID)
				if len(c.Tasks) == 0 {
					fmt.Println("      Тикеты: (0)")
					continue
				}
				for _, t := range c.Tasks {
					fmt.Printf("      ТИКЕТ: %s (ID: %s)\n", t.Title, t.ID)
				}
			}
		}
	}
}

/* ---------------- API calls ---------------- */

func getCompanies(domain, login, password string) ([]Company, error) {
	status, body, err := postJSON(
		fmt.Sprintf("https://%s/api-v2/auth/companies", domain),
		AuthRequest{Login: login, Password: password},
	)
	if err != nil {
		return nil, err
	}
	if status < 200 || status >= 300 {
		return nil, fmt.Errorf("status %d: %s", status, string(body))
	}

	var lr ListResponse[Company]
	if err := json.Unmarshal(body, &lr); err != nil {
		return nil, fmt.Errorf("companies unmarshal: %w; body=%s", err, string(body))
	}
	return lr.Content, nil
}

func getAPIKey(domain, login, password, companyID string) (string, error) {
	status, body, err := postJSON(
		fmt.Sprintf("https://%s/api-v2/auth/keys", domain),
		KeyRequest{Login: login, Password: password, CompanyID: companyID},
	)
	if err != nil {
		return "", err
	}
	if status < 200 || status >= 300 {
		return "", fmt.Errorf("status %d: %s", status, string(body))
	}

	var kr KeyResponse
	if err := json.Unmarshal(body, &kr); err != nil {
		return "", fmt.Errorf("key unmarshal: %w; body=%s", err, string(body))
	}
	if kr.Key == "" {
		return "", fmt.Errorf("пустой key: %s", string(body))
	}
	fmt.Println(kr.Key)
	return kr.Key, nil
}

func getProjects(domain, apiKey string) ([]Project, error) {
	status, body, err := getJSON(domain, "/projects", apiKey)
	if err != nil {
		return nil, err
	}
	if status < 200 || status >= 300 {
		return nil, fmt.Errorf("status %d: %s", status, string(body))
	}

	var lr ListResponse[Project]
	if err := json.Unmarshal(body, &lr); err != nil {
		return nil, fmt.Errorf("boards unmarshal: %w; body=%s", err, string(body))
	}
	return lr.Content, nil
}

func getBoards(domain, apiKey string) ([]Board, error) {
	status, body, err := getJSON(domain, "/boards", apiKey)
	if err != nil {
		return nil, err
	}
	if status < 200 || status >= 300 {
		return nil, fmt.Errorf("status %d: %s", status, string(body))
	}

	var lr ListResponse[Board]
	if err := json.Unmarshal(body, &lr); err != nil {
		return nil, fmt.Errorf("boards unmarshal: %w; body=%s", err, string(body))
	}
	return lr.Content, nil
}

// columns list
func getAllColumnIDs(domain, apiKey string) ([]string, error) {
	status, body, err := getJSON(domain, "/columns", apiKey)
	if err != nil {
		return nil, err
	}
	if status < 200 || status >= 300 {
		return nil, fmt.Errorf("status %d: %s", status, string(body))
	}

	var lr ListResponse[struct {
		ID string `json:"id"`
	}]
	if err := json.Unmarshal(body, &lr); err != nil {
		return nil, fmt.Errorf("columns list unmarshal: %w; body=%s", err, string(body))
	}

	ids := make([]string, 0, len(lr.Content))
	for _, c := range lr.Content {
		if c.ID != "" {
			ids = append(ids, c.ID)
		}
	}
	return ids, nil
}

// columns detail + mapping to board
func getColumnsByBoard(domain, apiKey string, columnIDs []string) (map[string][]ColumnDetail, error) {
	byBoard := map[string][]ColumnDetail{}
	for _, id := range columnIDs {
		col, err := getColumnByID(domain, apiKey, id)
		if err != nil {
			return nil, err
		}
		if col.BoardID == "" {
			return nil, fmt.Errorf("в колонке %s нет boardId: %+v", id, col)
		}
		byBoard[col.BoardID] = append(byBoard[col.BoardID], col)
	}
	return byBoard, nil
}

func getColumnByID(domain, apiKey, columnID string) (ColumnDetail, error) {
	status, body, err := getJSON(domain, "/columns/"+columnID, apiKey)
	if err != nil {
		return ColumnDetail{}, err
	}
	if status < 200 || status >= 300 {
		return ColumnDetail{}, fmt.Errorf("status %d: %s", status, string(body))
	}

	var col ColumnDetail
	if err := json.Unmarshal(body, &col); err != nil {
		return ColumnDetail{}, fmt.Errorf("column unmarshal: %w; body=%s", err, string(body))
	}
	return col, nil
}

func getTasksByColumn(domain, apiKey, columnID string) ([]Task, error) {
	q := url.Values{}
	q.Set("columnId", columnID)
	endpointA := "/tasks?" + q.Encode()

	status, body, err := getJSON(domain, endpointA, apiKey)
	if err == nil && status >= 200 && status < 300 {
		var lr ListResponse[Task]
		if err := json.Unmarshal(body, &lr); err == nil {
			return lr.Content, nil
		}
		// если формат другой — покажем ошибку парсинга
		return nil, fmt.Errorf("tasks unmarshal (A): %v; body=%s", err, string(body))
	}

	var lr ListResponse[Task]
	if err := json.Unmarshal(body, &lr); err != nil {
		return nil, fmt.Errorf("tasks unmarshal (B): %w; body=%s", err, string(body))
	}
	return lr.Content, nil
}

/* ---------------- HTTP helpers ---------------- */

func postJSON(url string, payload any) (status int, body []byte, err error) {
	b, _ := json.Marshal(payload)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	return resp.StatusCode, body, nil
}

func getJSON(domain, endpoint, apiKey string) (status int, body []byte, err error) {
	client := &http.Client{Timeout: 30 * time.Second}
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://%s/api-v2%s", domain, endpoint), nil)
	req.Header.Set("Authorization", "Bearer "+apiKey) // так требует дальнейшие запросы [web:34]
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	return resp.StatusCode, body, nil
}

/* ---------------- misc ---------------- */

func sortColumns(cols []ColumnDetail) {
	sort.SliceStable(cols, func(i, j int) bool {
		pi := columnSortKey(cols[i])
		pj := columnSortKey(cols[j])

		if pi != nil && pj != nil && *pi != *pj {
			return *pi < *pj
		}
		if pi != nil && pj == nil {
			return true
		}
		if pi == nil && pj != nil {
			return false
		}
		return strings.ToLower(cols[i].Title) < strings.ToLower(cols[j].Title)
	})
}

func columnSortKey(c ColumnDetail) *int {
	if c.Position != nil {
		return c.Position
	}
	if c.Index != nil {
		return c.Index
	}
	if c.Order != nil {
		return c.Order
	}
	return nil
}

func upsertEnvKey(filename, key, value string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		// если файла нет — создадим
		if os.IsNotExist(err) {
			return os.WriteFile(filename, []byte(fmt.Sprintf("%s=%s\n", key, value)), 0600)
		}
		return err
	}

	lines := strings.Split(string(data), "\n")
	found := false
	for i, line := range lines {
		if strings.HasPrefix(line, key+"=") {
			lines[i] = fmt.Sprintf("%s=%s", key, value)
			found = true
			break
		}
	}
	if !found {
		// вставим перед последней пустой строкой (если есть)
		if len(lines) > 0 && lines[len(lines)-1] == "" {
			lines = append(lines[:len(lines)-1], fmt.Sprintf("%s=%s", key, value), "")
		} else {
			lines = append(lines, fmt.Sprintf("%s=%s", key, value))
		}
	}

	out := strings.Join(lines, "\n")
	return os.WriteFile(filename, []byte(out), 0600)
}
