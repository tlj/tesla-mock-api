FROM golang:alpine as builder

RUN mkdir /build

ADD . /build/

WORKDIR /build

RUN go build -o tesla-mock-api cmd/mock/main.go

# PRODUCTION
FROM alpine

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY --from=builder /build/cmd/mock/fixtures /app/cmd/mock/fixtures
COPY --from=builder /build/tesla-mock-api /app

WORKDIR /app

ENTRYPOINT ["./tesla-mock-api"]
CMD [""]