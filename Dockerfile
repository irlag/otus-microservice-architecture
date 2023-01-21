ARG GOLANG_VERSION

FROM golang:${GOLANG_VERSION}-alpine AS app

WORKDIR /app

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ./

ARG VERSION
ENV CGO_ENABLED=0

RUN go build -ldflags "-s -X otus-microservice-architecture/cmd.Version=$VERSION" -v -a -o /bin/otus-microservice-architecture main.go

ENTRYPOINT ["/bin/otus-microservice-architecture"]