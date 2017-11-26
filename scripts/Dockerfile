FROM golang

ADD . github.com/tannergr/trivia
RUN go install github.com/tannergr/trivia
ENTRYPOINT ~/go/bin/basic_web_server

EXPOSE 3000
