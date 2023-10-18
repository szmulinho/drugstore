FROM golang:1.21.1-alpine AS build

WORKDIR /drugstore
COPY . .

RUN apk add --no-cache git
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o drugstore

FROM alpine:latest

WORKDIR /drugstore
COPY --from=build /drugstore/drugstore /drugstore/drugstore

COPY cert.pem /drugstore/cert.pem
COPY key.pem /drugstore/key.pem

EXPOSE 443

CMD ["./drugstore", "-cert", "/drugstore/cert.pem", "-key", "/drugstore/key.pem", "-addr", ":443"]
