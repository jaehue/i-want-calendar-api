FROM pangpanglabs/golang:builder AS builder

ADD . /go/src/github.com/jaehue/i-want-calendar-api
WORKDIR /go/src/github.com/jaehue/i-want-calendar-api
ENV CGO_ENABLED=0
RUN go build -o i-want-calendar-api

FROM pangpanglabs/alpine-ssl
WORKDIR /go/src/github.com/jaehue/i-want-calendar-api
COPY --from=builder /go/src/github.com/jaehue/i-want-calendar-api/i-want-calendar-api /go/src/github.com/jaehue/i-want-calendar-api/

EXPOSE 8000

CMD ["./i-want-calendar-api"]
