FROM golang:1.22.2

WORKDIR /app

EXPOSE 600

ARG SSH_PRIVATE_KEY
ARG SSH_PUBLIC_KEY

RUN apt-get update && apt-get install -y ca-certificates git-core ssh

RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/

RUN echo "$SSH_PRIVATE_KEY"
RUN echo "$SSH_PUBLIC_KEY"

RUN mkdir -p /root/.ssh && \
chmod 0700 /root/.ssh && \
ssh-keyscan github.com > /root/.ssh/known_hosts

RUN private.txt > /root/.ssh/id_ed25519 && \
chmod 600 /root/.ssh/id_ed25519
RUN public.txt > /root/.ssh/id_ed25519.pub && \
chmod 600 /root/.ssh/id_ed25519.pub


COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /proxy

CMD ["/proxy"]