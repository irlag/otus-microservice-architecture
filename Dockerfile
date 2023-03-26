ARG GOLANG_VERSION

FROM golang:${GOLANG_VERSION}-alpine AS build

WORKDIR /app

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ./

ARG VERSION
ENV CGO_ENABLED=0

RUN go build -ldflags "-s -X otus-microservice-architecture/cmd.Version=$VERSION" -v -a -o /bin/otus-microservice-architecture main.go

FROM scratch
COPY --from=build /bin/otus-microservice-architecture /bin/otus-microservice-architecture
COPY --from=build /app/migrations /migrations

ENTRYPOINT ["/bin/otus-microservice-architecture"]