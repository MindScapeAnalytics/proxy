FROM golang:1.22.2

WORKDIR /app

EXPOSE 6000

RUN GOCACHE=OFF

ARG USERNAME
ARG ACCESS_TOKEN
ENV ACCESS_TOKEN=$ACCESS_TOKEN

RUN git config --global url."https://${ACCESS_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

COPY . .
# -----------------------------------------
# ARG SSH_PRIVATE_KEY

# RUN apt-get update && apt-get install -y ca-certificates git-core ssh
# RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/

# RUN echo ${SSH_PRIVATE_KEY}

# RUN mkdir -p /root/.ssh && \
# chmod 0700 /root/.ssh && \
# ssh-keyscan github.com > /root/.ssh/known_hosts

# RUN echo ${SSH_PRIVATE_KEY} > /root/.ssh/id_ed25519 && \
# chmod 600 /root/.ssh/id_ed25519 && \
# # go mod download && \
# rm -rf /root/.ssh/


# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /proxy

CMD ["/proxy"]