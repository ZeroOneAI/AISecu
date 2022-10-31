# Build the manager binary
FROM golang:1.19 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY cmd/scanner cmd/scanner
COPY pkg pkg

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager cmd/scanner/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install github.com/google/go-containerregistry/cmd/crane@latest
RUN curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin v0.31.3
RUN apt-get update && apt-get install -y curl


# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
# FROM gcr.io/distroless/static:nonroot
FROM alpine
WORKDIR /
COPY --from=builder /workspace/manager .
COPY --from=builder /go/bin/crane /usr/local/bin/crane
COPY --from=builder /usr/local/bin/trivy /usr/local/bin/trivy
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/
# USER 65532:65532

ENTRYPOINT ["/manager"]
