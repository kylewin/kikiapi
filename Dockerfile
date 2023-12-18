FROM golang:1.21 AS build-stage
WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /kikiapi

FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /

COPY --from=build-stage /kikiapi /kikiapi
EXPOSE 8088
USER nonroot:nonroot
ENTRYPOINT ["/kikiapi"]
