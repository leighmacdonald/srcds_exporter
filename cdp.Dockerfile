FROM golang:1.16-alpine as build
LABEL maintainer="Leigh MacDonald <leigh.macdonald@gmail.com>"
WORKDIR /build
RUN apk add make git gcc libc-dev curl
ENV PATH="/go/bin:${PATH}"
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN PREFIX=. make

FROM chromedp/headless-shell:latest
LABEL maintainer="Leigh MacDonald <leigh.macdonald@gmail.com>"
RUN apt update
RUN apt install dumb-init -y
WORKDIR /app
COPY --from=build /build/srcds_exporter .
ENTRYPOINT ["dumb-init", "--"]
CMD ["./srcds_exporter", "-collectors.enabled", "map,players,rank", "-config.file", "/app/srcds.yaml"]
