MOCKGEN=mockgen
COVERAGE_FILE=coverage.out
MOCK_SRC_PLACES=internal/pkg/places/interfaces.go
MOCK_DST_PLACES=internal/pkg/places/mocks/mock.go
MOCK_SRC_AUTH=internal/pkg/auth/interfaces.go
MOCK_DST_AUTH=internal/pkg/auth/mocks/mock.go

all: test

mocks:
	$(MOCKGEN) -source=$(MOCK_SRC_PLACES) -destination=$(MOCK_DST_PLACES)
	$(MOCKGEN) -source=$(MOCK_SRC_AUTH) -destination=$(MOCK_DST_AUTH)

test: mocks
	go test places auth -coverprofile=$(COVERAGE_FILE)

cover: test
	go tool cover -html=$(COVERAGE_FILE)

clean:
	rm -f $(MOCK_DST)
	rm -f $(COVERAGE_FILE)

.PHONY: all mocks test cover clean
