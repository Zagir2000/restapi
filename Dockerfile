FROM golang:1.18.1

RUN mkdir -p /src/
WORKDIR /src/

COPY  . ./src/
COPY go.mod ./
COPY go.sum ./
RUN go mod download

WORKDIR /cmd/main/

CMD ["/cmd/main/", app.go]