FROM golang:1.16.0

RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

COPY custom_report.go main.go monitoring.go starwars.go /app/
RUN CGO_ENABLED=0 GOOS=linux GARCH=amd64 go build -o sample

EXPOSE 8000

CMD /app/sample