FROM golang:alpine AS build
WORKDIR /drugstore
LABEL maintainer="szmulinho"
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
FROM alpine:latest
WORKDIR /root/
COPY --from=build /drugstore .
EXPOSE 8093
CMD ["./drugstore"]

