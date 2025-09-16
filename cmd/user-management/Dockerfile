FROM golang:1.25.0 AS builder
ARG ARCH=amd64

WORKDIR /code

COPY go.mod go.sum vendor /code/
COPY . /code
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} go build \
		-mod=mod  \
		-o $GOPATH/bin/user-management \
		-gcflags "all=-trimpath=$GOPATH" \
		-a ./cmd/user-management

FROM alpine:3.15

COPY --from=builder /go/bin/user-management /bin/user-management
ENTRYPOINT ["/bin/user-management"]