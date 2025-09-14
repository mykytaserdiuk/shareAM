package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func getAPIKeyPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Не удалось получить домашнюю директорию:", err)
		os.Exit(1)
	}
	return filepath.Join(home, ".mycli_apikey")
}

func saveAPIKey(key string) error {
	path := getAPIKeyPath()
	return os.WriteFile(path, []byte(key), 0600) // только для пользователя
}

func loadAPIKey() (string, error) {
	path := getAPIKeyPath()
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func main() {
	// Флаги CLI
	filePath := flag.String("file", "", "Путь к файлу для загрузки")
	apiURL := flag.String("url", "", "URL API для загрузки файла")
	setKey := flag.Bool("setkey", false, "Сохранить API ключ")
	flag.Parse()

	if *setKey {
		fmt.Print("Введите ваш API ключ: ")
		reader := bufio.NewReader(os.Stdin)
		key, _ := reader.ReadString('\n')
		key = strings.TrimSpace(key)
		err := saveAPIKey(key)
		if err != nil {
			fmt.Println("Ошибка при сохранении ключа:", err)
			return
		}
		fmt.Println("API ключ сохранен!")
		return
	}

	fmt.Println(*filePath)
	fmt.Println(*apiURL)
	if *filePath == "" || *apiURL == "" {
		fmt.Println("Использование: go run main.go -file=путь_к_файлу -url=адрес_API")
		fmt.Println("Для сохранения ключа: go run main.go -setkey")
		return
	}

	apiKey, err := loadAPIKey()
	if err != nil {
		fmt.Println("Не найден API ключ. Сначала сохраните его с помощью -setkey")
		return
	}

	// Открываем файл
	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	// Создаем multipart/form-data запрос
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		fmt.Println("Ошибка при создании формы:", err)
		return
	}
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println("Ошибка при копировании файла:", err)
		return
	}
	writer.Close()

	req, err := http.NewRequest("POST", *apiURL, body)
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("User-Token", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	fmt.Println("Ответ сервера:", string(respBody))
}
