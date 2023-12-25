FROM golang:1.20.5-buster as builder
WORKDIR /splunk-cf-logdrain
COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$GOARCH go build -o /splunk-cf-logdrain

## Build final image
FROM fluent/fluent-bit:2.2.1-debug
LABEL maintainer="andy.lo-a-foe@philips.com"
RUN apt-get update && apt-get install -y \
    jq \
    supervisor \
    curl \
 && rm -rf /var/lib/apt/lists/*

RUN mkdir -p /sidecars/bin /sidecars/supervisor/conf.d sidecars/etc

COPY --from=builder     /splunk-cf-logdrain        /sidecars/bin

COPY docker/supervisord_configs/ /sidecars/supervisor/conf.d
COPY docker/config/fluent-bit.conf /fluent-bit/etc/fluent-bit.conf
COPY docker/config/parsers.conf /fluent-bit/etc/parsers.conf

EXPOSE 8080

COPY docker/supervisord.conf /etc/
CMD ["supervisord", "--nodaemon", "--configuration", "/etc/supervisord.conf"]
