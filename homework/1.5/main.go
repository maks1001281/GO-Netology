package main

import (
        "errors"
        "fmt"
        "os"
)

func main() {
        content := ReadProcessWrite()
        process(content)
}

//Читаем файл, обрабатываем ошибки
func ReadProcessWrite() string {
        content, err := os.ReadFile("file/task.yaml")
        if err != nil {
                if errors.Is(err, os.ErrNotExist) {
                        panic("Файл не найден!")
                } else if errors.Is(err, os.ErrPermission) {
                        panic("Ошибка доступа")
                } else if errors.Is(err, os.ErrInvalid) {
                        panic("Файл не текстовый")
                } else if errors.Is(err, os.ErrDeadlineExceeded) {
                        panic("Время ожидания истекло")
                }
        }
        fmt.Println("Содержимое:\n", string(content))
        return string(content)
}

//Создаем файл если его нет и записывает input с ReadProcessWrite в taskoutput.yaml
func process(content string) {
        if err := os.WriteFile("file/taskoutput.yaml", []byte(content), 0644); err != nil {
                if errors.Is(err, os.ErrPermission) {
                        fmt.Println("Ошибка доступа:")
                }
        }
}