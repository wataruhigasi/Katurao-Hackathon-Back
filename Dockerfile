# syntax=docker/dockerfile:1

FROM golang:1.21 AS build-stage

WORKDIR /app

COPY src/go.mod src/go.sum ./

RUN go mod download

COPY ./src ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /katurao-hackathon-back


# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /katurao-hackathon-back /katurao-hackathon-back

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/katurao-hackathon-back"]


# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
# EXPOSE 8080

# CMD ["/katurao-hackathon-back"]
