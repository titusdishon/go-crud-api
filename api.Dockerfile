FROM golang:1.19-alpine

WORKDIR /app

RUN apk update && \
    apk add libc-dev && \
    apk add gcc && \
    apk add make

COPY ./go.mod go.sum ./
RUN go mod download && go mod verify


# Install Compile Daemon for go. We'll use it to watch changes in go files
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

COPY . .
COPY ./entrypoint.sh /entrypoint.sh

# wait-for-it requires bash, which alpine doesn't ship with by default. Use wait-for instead
ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for
RUN chmod +rx /usr/local/bin/wait-for /entrypoint.sh

ENTRYPOINT [ "sh", "/entrypoint.sh" ]
