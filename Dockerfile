# syntax=docker/dockerfile:1
FROM golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY cmd ./cmd
COPY internal ./internal
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/colour ./cmd/colour

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/colour /colour
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/colour"]
