FROM golang:1.16.0 as builder

RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

COPY custom_report.go main.go monitoring.go starwars.go /app/
RUN CGO_ENABLED=0 GOOS=linux GARCH=amd64 go build -a -ldflags '-extldflags "-static" -s -w' -o sample

EXPOSE 8000

CMD /app/sample

############################################################

FROM alpine

COPY --from=builder /app/sample /app/sample

EXPOSE 8000
CMD ["/app/sample"]

