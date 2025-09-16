FROM golang:1.25.0 AS builder
ARG ARCH=amd64

WORKDIR /code

COPY go.mod go.sum vendor /code/
COPY . /code
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} go build \
		-mod=mod  \
		-o $GOPATH/bin/file-storage \
		-gcflags "all=-trimpath=$GOPATH" \
		-a ./cmd/file-storage

FROM alpine:3.15

COPY --from=builder /go/bin/file-storage /bin/file-storage
ENTRYPOINT ["/bin/file-storage"]