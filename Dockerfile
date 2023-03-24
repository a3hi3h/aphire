FROM golang:alpine AS build
WORKDIR /go/src/app
COPY . .
RUN go mod tidy
RUN go build -o ./bin/web-app ./main.go
COPY .env ./bin

FROM alpine:3.17
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin
EXPOSE 80
ENTRYPOINT /go/bin/web-app --port 80