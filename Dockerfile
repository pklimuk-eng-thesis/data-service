FROM golang:1.20.1-alpine AS builder
LABEL maintainer="Pavel Klimuk <pavelklimuk@outlook.com>"

WORKDIR /src 
COPY . .
RUN go mod download
RUN go build -o dataService cmd/main.go


FROM scratch
WORKDIR /
COPY --from=builder /src/dataService /dataService
EXPOSE 8087
CMD ["/dataService"]