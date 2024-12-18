FROM golang:1.23.1-alpine AS builder
COPY . /github.com/go-park-mail-ru/2024_2_ThereWillBeName/survey
WORKDIR /github.com/go-park-mail-ru/2024_2_ThereWillBeName/survey
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -o ./.bin ./cmd/survey/main.go
FROM scratch AS runner
WORKDIR /build
COPY --from=builder /github.com/go-park-mail-ru/2024_2_ThereWillBeName/survey/.bin .
COPY --from=builder /github.com/go-park-mail-ru/2024_2_ThereWillBeName/survey/config config/
EXPOSE 50054
ENTRYPOINT ["./.bin"]
