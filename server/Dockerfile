FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -o shareterm_server .

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/shareterm_server /app/
COPY --from=builder /build/public /app/public
VOLUME ["/app/data"]
WORKDIR /app
CMD ["./shareterm_server"]