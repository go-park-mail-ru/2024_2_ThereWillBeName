FROM golang:1.23.1-alpine AS builder
COPY . /github.com/go-park-mail-ru/2024_2_ThereWillBeName/attractions
WORKDIR /github.com/go-park-mail-ru/2024_2_ThereWillBeName/attractions
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -o ./.bin ./cmd/attractions/main.go
FROM scratch AS runner
WORKDIR /build
COPY --from=builder /github.com/go-park-mail-ru/2024_2_ThereWillBeName/attractions/.bin .
COPY --from=builder /github.com/go-park-mail-ru/2024_2_ThereWillBeName/attractions/config config/
EXPOSE 50051 8091
ENTRYPOINT ["./.bin"]
