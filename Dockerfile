FROM golang:1.21.1-alpine AS build

WORKDIR /drugstore
COPY . .

RUN apk add --no-cache git
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o drugstore

FROM alpine:latest

WORKDIR /drugstore
COPY --from=build /drugstore/drugstore /drugstore/drugstore

EXPOSE 8091

CMD ["./drugstore"]
