FROM golang

RUN go get github.com/gorilla/mux

RUN go get github.com/adlio/trello

WORKDIR /go/src/github.com/heaptracetechnology/microservice-trello

ADD . /go/src/github.com/heaptracetechnology/microservice-trello

RUN go install github.com/heaptracetechnology/microservice-trello

ENTRYPOINT microservice-trello

EXPOSE 3000