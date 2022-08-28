FROM golang:latest AS builder
WORKDIR /app
COPY go.* .
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -o remo-exporter . 

FROM alpine:3.9
COPY --from=builder /app/remo-exporter ./remo-exporter

RUN apk --no-cache add ca-certificates \
     && addgroup exporter \
     && adduser -S -G exporter exporter
USER exporter

ENV LISTEN_PORT=9352
EXPOSE 9352
ENTRYPOINT [ "./remo-exporter" ]