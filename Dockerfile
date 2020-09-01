# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang
ADD . /go/src/github.com/Andrew-Klaas/go-movies-favoritess-app
WORKDIR /go/src/github.com/Andrew-Klaas/go-movies-favoritess-app
RUN go install /go/src/github.com/Andrew-Klaas/go-movies-favoritess-app

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/go-movies-favoritess-app

# Document that the service listens on port 8080.
EXPOSE 8081


#docker build --no-cache -t aklaas2/go-movies-favorites-app .;docker push aklaas2/go-movies-favorites-app:latest