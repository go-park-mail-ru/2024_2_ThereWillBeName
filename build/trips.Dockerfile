FROM golang:1.23.1-alpine AS builder
COPY . /github.com/go-park-mail-ru/2024_2_ThereWillBeName/trips
WORKDIR /github.com/go-park-mail-ru/2024_2_ThereWillBeName/trips
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -o ./.bin ./cmd/trips/main.go
FROM scratch AS runner
WORKDIR /build
COPY --from=builder /github.com/go-park-mail-ru/2024_2_ThereWillBeName/trips/.bin .
COPY --from=builder /github.com/go-park-mail-ru/2024_2_ThereWillBeName/trips/config config/
EXPOSE 50053
ENTRYPOINT ["./.bin"]
