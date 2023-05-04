FROM golang:1.20-bullseye AS build

WORKDIR /hello
COPY hello/* ./
RUN go mod download
RUN go build -o /hello_http

FROM gcr.io/distroless/static-debian11
LABEL org.opencontainers.image.title "Hello from Path"
LABEL org.opencontainers.image.licenses "Apache-2.0 License"
LABEL org.opencontainers.inage.description "Hello from Path illustrates the build of a small Go HTTP application into a container."

COPY --from=build /hello_http /bin/

EXPOSE 8080
ENTRYPOINT ["/bin/hello_http"]