FROM golang:1.21-alpine as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/server .

RUN touch .env

EXPOSE 8000

CMD [ "./server" ]