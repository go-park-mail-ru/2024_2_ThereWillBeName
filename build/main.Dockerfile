FROM golang:1.22.7-alpine AS builder
COPY . /github.com/go-park-mail-ru/2024_2_ThereWillBeName
WORKDIR /github.com/go-park-mail-ru/2024_2_ThereWillBeName
RUN go mod download
RUN go clean --modcache
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -o ./.bin ./cmd/main/main.go
FROM scratch AS runner
WORKDIR /build
COPY --from=builder /github.com/go-park-mail-ru/2024_2_ThereWillBeName/.bin .
COPY --from=builder /github.com/go-park-mail-ru/2024_2_ThereWillBeName/config config/
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
ENV TZ="Europe/Moscow"
ENV ZONEINFO=/zoneinfo.zip
EXPOSE 80 443 8080
ENTRYPOINT ["./.bin"]