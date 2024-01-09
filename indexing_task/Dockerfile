FROM golang:1.22rc1

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /indexing

ENTRYPOINT ["/indexing"]
