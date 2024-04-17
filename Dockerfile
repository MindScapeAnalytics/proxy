FROM golang:1.22.2

WORKDIR /app

EXPOSE 6000

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /proxy

CMD ["/proxy"]