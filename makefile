# Определите переменные
PKG=./...
MOCKGEN=mockgen
COVERAGE_FILE=coverage.out
MOCK_SRC=internal/pkg/auth/interfaces.go
MOCK_DST=internal/pkg/auth/mocks/mock.go
PACKAGE_NAME=auth

# По умолчанию - запуск всех тестов
all: test

# Генерация моков, если исходные файлы изменились
mocks: $(MOCK_DST)
	@echo "Generating mocks..."
	$(MOCKGEN) -source=$(MOCK_SRC) -destination=$(MOCK_DST) -package=$(PACKAGE_NAME)

# Запуск тестов с покрытием
test: mocks
	@echo "Running tests with coverage..."
	go test $(PKG) -coverprofile=$(COVERAGE_FILE)

# Просмотр отчета о покрытии
cover: test
	@echo "Generating coverage report..."
	go tool cover -html=$(COVERAGE_FILE)

# Очистка сгенерированных файлов
clean:
	@echo "Cleaning up..."
	rm -f $(MOCK_DST)
	rm -f $(COVERAGE_FILE)

.PHONY: all mocks test cover clean