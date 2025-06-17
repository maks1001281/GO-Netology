# Имя исполняемого файла
BINARY_NAME=hw

# Директория с bin
SRC_DIR=./build

# Цель по умолчанию
all: build

# Сборка bin приложения
build:
	go build -o $(SRC_DIR)/$(BINARY_NAME) *.go

# Запуск bin приложения
run: build
	$(SRC_DIR)/$(BINARY_NAME)

# Очистка скомпилированных файлов
clean:
	rm -f $(SRC_DIR)/$(BINARY_NAME)

# Показать помощь
help:
	@echo "Makefile для моих go проектов"
	@echo "Доступные команды:"
	@echo "  make build  - Сборка приложения"
	@echo "  make run    - Сборка и запуск приложения"
	@echo "  make clean  - Очистка скомпилированных файлов"
	@echo "  make help   - Показать эту помощь"

.PHONY: all build run clean help