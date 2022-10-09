FROM --platform=$BUILDPLATFORM golang:1.19.1-alpine3.16 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download -x

COPY main.go ./
# Purposefully set AFTER downloading and caching dependencies
ARG TARGETOS TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-w -s" -o /go/bin/github-ratelimit-metrics .

FROM gcr.io/distroless/static-debian11:nonroot
COPY --from=builder /go/bin/github-ratelimit-metrics /usr/local/bin/github-ratelimit-metrics

EXPOSE 8000
ENTRYPOINT ["github-ratelimit-metrics"]
