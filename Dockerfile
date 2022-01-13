FROM golang:1.15-alpine AS build

ARG VERSION

LABEL description="An incredibly fast proxy checker & IP rotator with ease."
LABEL repository="https://github.com/kitabisa/mubeng"
LABEL maintainer="dwisiswant0"

WORKDIR /app
COPY ./go.mod .
RUN go mod download

COPY . .
RUN go build -ldflags "-s -w -X ktbs.dev/mubeng/common.Version=${VERSION}" \
	-o ./bin/mubeng ./cmd/mubeng 

FROM alpine:latest

COPY --from=build /app/bin/mubeng /bin/mubeng
ENV HOME /
COPY ./proxies.txt /proxies.txt
COPY ./proxies_ipv6.txt /proxies_ipv6.txt

EXPOSE 23528 23529
CMD ["/bin/mubeng", "-a", "127.0.0.1:23528", "-f", "/proxies.txt"]
# CMD ["/bin/mubeng", "-a", "127.0.0.1:23529", "-f", "/proxies_ipv6.txt"]
