FROM golang:1.19

ENV APP_HOME /go/src/transaction-risk

WORKDIR "$APP_HOME"
COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o transaction-risk

EXPOSE 3000
CMD ["./transaction-risk"]