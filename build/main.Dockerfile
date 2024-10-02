FROM golang:1.23.1-alpine AS builder
COPY . /github.com/go-park-mail-ru/2024_2_ThereWillBeName
WORKDIR /github.com/go-park-mail-ru/2024_2_ThereWillBeName
RUN go mod download
RUN go clean --modcache
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -o ./.bin ./cmd/main.go
FROM scratch AS runner
WORKDIR /build
COPY --from=builder /github.com/go-park-mail-ru/2024_2_ThereWillBeName/.bin .
EXPOSE 8080
ENTRYPOINT ["./.bin"]