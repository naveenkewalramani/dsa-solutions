FROM golang:latest

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY main.go ./

RUN go build -o /data_structure_impl

CMD ["/data_structure_impl"]
