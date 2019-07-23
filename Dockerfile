FROM golang
WORKDIR /go
COPY . .
ENV GOPATH="/go"
ENTRYPOINT [ "go install" ]