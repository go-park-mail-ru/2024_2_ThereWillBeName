FROM golang:1.23.1-alpine AS builder
COPY . /github.com/go-park-mail-ru/2024_2_ThereWillBeName/gateway
WORKDIR /github.com/go-park-mail-ru/2024_2_ThereWillBeName/gateway
RUN go mod download
RUN go clean --modcache
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -o ./.bin ./cmd/gateway/main.go
FROM scratch AS runner
WORKDIR /build
COPY --from=builder /github.com/go-park-mail-ru/2024_2_ThereWillBeName/gateway/.bin .
COPY --from=builder /github.com/go-park-mail-ru/2024_2_ThereWillBeName/gateway/config config/
EXPOSE 8081
ENTRYPOINT ["./.bin"]
