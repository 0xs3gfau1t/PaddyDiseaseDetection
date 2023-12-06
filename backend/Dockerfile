FROM golang:1.21

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum docker-compose-prepare.sh ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v

ENTRYPOINT [ "/app/docker-compose-prepare.sh" ]
