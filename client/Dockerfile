# Builder stage
FROM golang:alpine3.14 AS builder

# Install dependencies
RUN apk update && apk add --no-cache \
  git \
  ca-certificates \
  && update-ca-certificates

# Add source files and set the proper workdir |
COPY . $GOPATH/src/github.com/josedelri85/faceit_technical_test/client/
WORKDIR $GOPATH/src/github.com/josedelri85/faceit_technical_test/client/

# Enable GO Modules
ENV GO111MODULES=on
# Build the binary
RUN go build -mod=vendor -o /go/bin/faceit_technical_test/client .

# Final image
FROM alpine

# Copy our static executable
COPY --from=builder /go/bin/faceit_technical_test/client /go/bin/faceit_technical_test/client

# Copy the ca-certificates to be able to perform https requests
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs

RUN apk add --no-cache tzdata && apk add curl

# Run the binary
ENTRYPOINT ["/go/bin/faceit_technical_test/client"]