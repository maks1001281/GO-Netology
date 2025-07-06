package main

import (
        "errors"
        "fmt"
        "io"
        "os"
        "strings"
)

func ReadProcessWrite(
        inputPath string,
        outputPath string,
        process func(string) (string, error),
) error {
        // Открываем входной файл
        inputFile, err := os.Open(inputPath)
        if err != nil {
          if errors.Is(err, os.ErrNotExist) {
            return fmt.Errorf("файл не найден:%w", err)
          }
          if errors.Is(err, os.ErrPermission) {
            return fmt.Errorf("ошибка доступа:%w", err)
          }
          if errors.Is(err, os.ErrInvalid) {
            return fmt.Errorf("некорректные данные:%w", err)
          }
          if errors.Is(err, os.ErrDeadlineExceeded) {
            return fmt.Errorf("время ожидания открытия файло истекло:%w", err)
          }
        }
        defer inputFile.Close()

        // Читаем содержимое файла
        inputData, err := io.ReadAll(inputFile)
        if err != nil {
          if errors.Is(err, os.ErrNotExist) {
            return fmt.Errorf("файл не найден:%w", err)
          }
          if errors.Is(err, os.ErrPermission) {
            return fmt.Errorf("ошибка доступа:%w", err)
          }
          if errors.Is(err, os.ErrInvalid) {
            return fmt.Errorf("некорректные данные:%w", err)
          }
          if errors.Is(err, os.ErrDeadlineExceeded) {
            return fmt.Errorf("время ожидания открытия файло истекло:%w", err)
          }
        }

        // Обрабатываем данные
        processedData, err := process(string(inputData))
        if err != nil {
          if errors.Is(err, os.ErrNotExist) {
            return fmt.Errorf("файл не найден:%w", err)
          }
          if errors.Is(err, os.ErrPermission) {
            return fmt.Errorf("ошибка доступа:%w", err)
          }
          if errors.Is(err, os.ErrInvalid) {
            return fmt.Errorf("некорректные данные:%w", err)
          }
          if errors.Is(err, os.ErrDeadlineExceeded) {
            return fmt.Errorf("время ожидания открытия файло истекло:%w", err)
          }
        }

        // Создаем/перезаписываем выходной файл
        outputFile, err := os.Create(outputPath)
        if err != nil {
          if errors.Is(err, os.ErrNotExist) {
            return fmt.Errorf("файл не найден:%w", err)
          }
          if errors.Is(err, os.ErrPermission) {
            return fmt.Errorf("ошибка доступа:%w", err)
          }
          if errors.Is(err, os.ErrInvalid) {
            return fmt.Errorf("некорректные данные:%w", err)
          }
          if errors.Is(err, os.ErrDeadlineExceeded) {
            return fmt.Errorf("время ожидания открытия файло истекло:%w", err)
          }
        }
        defer outputFile.Close()

        // Записываем результат
        _, err = outputFile.WriteString(processedData)
        if err != nil {
          if errors.Is(err, os.ErrNotExist) {
            return fmt.Errorf("файл не найден:%w", err)
          }
          if errors.Is(err, os.ErrPermission) {
            return fmt.Errorf("ошибка доступа:%w", err)
          }
          if errors.Is(err, os.ErrInvalid) {
            return fmt.Errorf("некорректные данные:%w", err)
          }
          if errors.Is(err, os.ErrDeadlineExceeded) {
            return fmt.Errorf("время ожидания открытия файло истекло:%w", err)
          }
        }

        return nil
}

func ToUpper(inputData string) (string, error) {
        return strings.ToUpper(inputData), nil
}

func main() {
        var inputPath string = "file/task.yaml"
        var outputPath string = "file/taskoutput.yaml"
        var oshibka =  ReadProcessWrite(inputPath, outputPath, ToUpper)
        fmt.Print(oshibka)
}