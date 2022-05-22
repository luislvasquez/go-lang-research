FROM golang:1.17

RUN mkdir /app
WORKDIR /app
ADD src /app

RUN go install github.com/githubnemo/CompileDaemon@latest

# Needs go.mod file (go mod init main + go mod tidy)
RUN go get -u
ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
