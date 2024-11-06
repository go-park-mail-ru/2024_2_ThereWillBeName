PKG=./...
MOCKGEN=mockgen
COVERAGE_FILE=coverage.out
MOCK_SRC_PLACES=internal/pkg/places/interfaces.go
MOCK_DST_PLACES=internal/pkg/places/mocks/mock.go
MOCK_SRC_USER=internal/pkg/user/interfaces.go
MOCK_DST_USER=internal/pkg/user/mocks/mock.go
PACKAGE_NAME_USER=user
PACKAGE_NAME_PLACES=places


all: test

mocks:
	$(MOCKGEN) -source=$(MOCK_SRC_PLACES) -destination=$(MOCK_DST_PLACES) -package=$(PACKAGE_NAME_PLACES)
	$(MOCKGEN) -source=$(MOCK_SRC_USER) -destination=$(MOCK_DST_USER) -package=$(PACKAGE_NAME_USER)

test: mocks
	go test $(PKG) -coverprofile=$(COVERAGE_FILE)

cover: test
	cat $(COVERAGE_FILE) | grep -v '_mock.go'
	go tool cover -func=$(COVERAGE_FILE) | grep total
	go tool cover -html=$(COVERAGE_FILE)

clean:
	rm -f $(MOCK_DST_PLACES)
	rm -f $(MOCK_DST_USER)
	rm -f $(COVERAGE_FILE)

.PHONY: all mocks test cover clean

build_:
	go build -o ./.bin cmd/main/main.go

run: build_
	./.bin

.PHONY: lint
lint:
	golangci-lint run --config=.golangci.yaml


