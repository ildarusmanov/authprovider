# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.8

# Copy the local package files to the container's workspace.
COPY . /go/src/github.com/ildarusmanov/authprovider

# setup dependencies
WORKDIR /go/src/github.com/ildarusmanov/authprovider
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure


RUN go install gituhb.com/ildarusmanov/authprovider

# Run the command by default when the container starts.
ENTRYPOINT /go/bin/eventmapper

# Document that the service listens on port 8000.
EXPOSE 3333