FROM --platform=$BUILDPLATFORM golang:alpine AS builder
ARG TARGETOS TARGETARCH
WORKDIR /build

ENV CGO_ENABLED 0
ENV GOOS $TARGETOS
ENV GOARCH $TARGETARCH

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY conf/ conf/
COPY pkg/ pkg/
COPY statik/ statik/
COPY main.go main.go
# Removed the debug message -ldflags="-s -w" to reduce the image size
RUN go build -ldflags="-s -w" -o mebius ./main.go

FROM --platform=$TARGETPLATFORM alpine
RUN apk update --no-cache && apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /build/mebius /app/mebius
ENTRYPOINT ["/app/mebius"]