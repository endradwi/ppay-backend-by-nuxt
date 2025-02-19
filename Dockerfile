FROM golang

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o backend main.go

ENTRYPOINT [ "/app/backend" ]