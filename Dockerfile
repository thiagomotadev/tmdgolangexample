FROM golang:1.15-buster

RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/postgres
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/thiagomotadev/tmdgolangbase

WORKDIR /go/src/github.com/thiagomotadev/tmdgolangexample

ENTRYPOINT [ "go", "run", "." ]
