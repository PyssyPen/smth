package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var actions = []string{
	"logged in",
	"logged out",
	"create record",
	"delete record",
	"update record",
}

type User struct {
	id    int
	email string
	logs  []logItem
}

type logItem struct {
	action    string
	timestamp time.Time
}

func (u User) getActivityInfo() string {
	out := fmt.Sprintf("ID: %d | Email: %s\nActivity Log:\n", u.id, u.email)
	for i, item := range u.logs {
		out += fmt.Sprintf("%d. [%s] at %s\n", i, item.action, item.timestamp)
	}

	return out

}

func generateUser(count int) []User {
	user := make([]User, count)

	for i := 0; i < count; i++ {
		user[i] = User{
			id:    i + 1,
			email: fmt.Sprintf("user%d@gmail", i+1),
			logs:  generateLogs(500 + rand.Intn(500)),
		}
	}

	return user
}

func generateLogs(count int) []logItem {
	logs := make([]logItem, count)

	for i := 0; i < count; i++ {
		logs[i] = logItem{
			timestamp: time.Now(),
			action:    actions[rand.Intn(len(actions)-1)],
		}
	}
	return logs
}

func saveUserInfo(user User) error {
	fmt.Printf("writting file for user id: %d\n", user.id)

	filename := fmt.Sprintf("logs/uid_%d.txt", user.id)

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644) // почитать про permission 0644 (read/write)
	if err != nil {
		return err
	}

	_, err = file.WriteString(user.getActivityInfo())
	return err
}

func main() {
	// для того, чтобы записать логи, надо сделать папку /VSC/GO/learn.udemy/2/logs
	rand.Seed(time.Now().Unix())

	users := generateUser(100)

	for _, user := range users {
		saveUserInfo(user)
	}
}
