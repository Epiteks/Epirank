# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/Shakarang/Epirank

# # Build the outyet command inside the container.
# # (You may fetch or manage dependencies here,
# # either manually or with a tool like "godep".)

RUN go get -u github.com/Sirupsen/logrus
RUN go get -u github.com/mattn/go-sqlite3

RUN go install github.com/Shakarang/Epirank

# # Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/Epirank

# Document that the service listens on port 8080.
#EXPOSE 8080E
CMD ["/go/bin/Epirank"]