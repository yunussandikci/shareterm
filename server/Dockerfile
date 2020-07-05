FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go build -o shareterm .

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/shareterm /app/
COPY --from=builder /build/public /app/public
VOLUME ["/app/data"]
WORKDIR /app
CMD ["./shareterm"]