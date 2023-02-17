FROM golang:1.19 as builder

WORKDIR /go/src

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /go/bin/main main.go

FROM scratch

COPY --from=builder /go/bin/main .
COPY --from=builder /go/src/.env .

ENTRYPOINT ["./main"]

