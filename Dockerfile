FROM golang:1.22.2

WORKDIR /app

EXPOSE 6000

RUN GOCACHE=OFF

ARG USERNAME
ARG ACCESS_TOKEN
ENV ACCESS_TOKEN=$ACCESS_TOKEN

RUN git config --global url."https://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:x-oauth-basic@github.com/".insteadOf "https://github.com/"

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /proxy

CMD ["/proxy"]