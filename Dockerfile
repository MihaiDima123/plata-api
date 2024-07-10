FROM golang:1.22.5

WORKDIR /app

COPY go.mod ./

RUN go mod download && go mod verify

COPY . /app

RUN go build -o plata cmd/main.go

CMD [ "/app/plata" ]
