FROM golang:1.23.1-alpine AS builder
COPY . /github.com/go-park-mail-ru/2024_2_ThereWillBeName/users
WORKDIR /github.com/go-park-mail-ru/2024_2_ThereWillBeName/users
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -o ./.bin ./cmd/users/main.go
FROM scratch AS runner
WORKDIR /build
COPY --from=builder /github.com/go-park-mail-ru/2024_2_ThereWillBeName/users/.bin .
COPY --from=builder /github.com/go-park-mail-ru/2024_2_ThereWillBeName/users/config config/
EXPOSE 50052
ENTRYPOINT ["./.bin"]
