FROM golang:1.22.2

WORKDIR /app

EXPOSE 600

ARG SSH_PRIVATE_KEY

RUN mkdir -p /root/.ssh && \
chmod 0700 /root/.ssh && \
ssh-keyscan github.com > /root/.ssh/known_hosts

RUN echo "$SSH_PRIVATE_KEY" > /root/.ssh/id_ed25519 && \
chmod 600 /root/.ssh/id_ed25519

RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /proxy

CMD ["/proxy"]