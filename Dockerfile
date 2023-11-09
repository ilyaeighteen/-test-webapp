FROM golang:1.21.3-alpine as build
WORKDIR /build
#устанавилваем необходимые go библиотеки
COPY . .
RUN go mod download
RUN go build -o /app cmd/test-application/main.go

#копируем собранное приложение в новый образ
FROM alpine
COPY --from=build app /bin/app
COPY cmd/test-application/templates /templates
ENTRYPOINT ["/bin/app"]
