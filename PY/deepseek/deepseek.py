##         python3 deepseek.py
import requests
import json


## КЛЮЧ ДОЛЖЕН ЛЕЖАТЬ В ТОЙ ЖЕ ПАПКЕ, ЧТО И ЭТА ПРОГРАММА И НАЗЫВАТЬСЯ "key.txt"

## это от прошлого владельца
## API_KEY = "" # внутри скобок свой апи ключ отсюда https://openrouter.ai/settings/keys
MODEL = "deepseek/deepseek-r1:free"



def read_api_key(file_path):
    try:
        with open(file_path, 'r') as file:
            # Читаем только первую строку
            api_key = file.readline().strip()
            return api_key
    except FileNotFoundError:
        print(f"Файл {file_path} не найден.")
        return None

# Укажите путь к файлу с API-ключом
api_key_file = '/home/pyssy/VSC/PY/deepseek/key.txt'

# Прочитайте API-ключ
api_key = read_api_key(api_key_file)

if api_key:
    print(f"Ваш API-ключ: {api_key}")







def process_content(content):
    return content.replace('<think>', '').replace('</think>', '')

def chat_stream(prompt):
    headers = {
        "Authorization": f"Bearer {api_key}",
        "Content-Type": "application/json"
    }
    
    data = {
        "model": MODEL,
        "messages": [{"role": "user", "content": prompt}],
        "stream": True
    }

    with requests.post(
        "https://openrouter.ai/api/v1/chat/completions",
        headers=headers,
        json=data,
        stream=True
    ) as response:
        if response.status_code != 200:
            print("Ошибка API:", response.status_code)
            return ""

        full_response = []
        
        for chunk in response.iter_lines():
            if chunk:
                chunk_str = chunk.decode('utf-8').replace('data: ', '')
                try:
                    chunk_json = json.loads(chunk_str)
                    if "choices" in chunk_json:
                        content = chunk_json["choices"][0]["delta"].get("content", "")
                        if content:
                            cleaned = process_content(content)
                            print(cleaned, end='', flush=True)
                            full_response.append(cleaned)
                except:
                    pass

        print()  # Перенос строки после завершения потока
        return ''.join(full_response)
def main():
    print("Чат с DeepSeek-R1\nДля выхода введите 'exit'\n")

    while True:
        user_input = input("Вы: ")
        
        if user_input.lower() == 'exit':
            print("Завершение работы...")
            break
            
        print("DeepSeek-R1:", end=' ', flush=True)
        chat_stream(user_input)

if __name__ == "__main__":
    main()