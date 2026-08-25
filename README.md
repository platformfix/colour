# colour

[![ci](https://github.com/platformfix/colour/actions/workflows/ci.yml/badge.svg)](https://github.com/platformfix/colour/actions/workflows/ci.yml)
[![e2e](https://github.com/platformfix/colour/actions/workflows/e2e.yml/badge.svg)](https://github.com/platformfix/colour/actions/workflows/e2e.yml)
[![lint](https://github.com/platformfix/colour/actions/workflows/lint.yml/badge.svg)](https://github.com/platformfix/colour/actions/workflows/lint.yml)
[![k8s-validate](https://github.com/platformfix/colour/actions/workflows/k8s-validate.yml/badge.svg)](https://github.com/platformfix/colour/actions/workflows/k8s-validate.yml)
[![commit-lint](https://github.com/platformfix/colour/actions/workflows/commit-lint.yaml/badge.svg)](https://github.com/platformfix/colour/actions/workflows/commit-lint.yaml)
[![pr-lint](https://github.com/platformfix/colour/actions/workflows/pr-lint.yml/badge.svg)](https://github.com/platformfix/colour/actions/workflows/pr-lint.yml)
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/platformfix/colour/badge)](https://scorecard.dev/viewer/?uri=github.com/platformfix/colour)
[![Latest Release](https://img.shields.io/github/v/release/platformfix/colour)](https://github.com/platformfix/colour/releases)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A tiny web server for literal blue/green Kubernetes deployment demos: the
page's background colour comes from the pod's hostname, so a Deployment
named `blue` serves a blue page and one named `green` serves a green page.

Inspired by [jpetazzo/color](https://github.com/jpetazzo/color), rebuilt for
Platform Fix's own Kubernetes workshops.

## Quickstart

Run the raw demo (the whole point of this tool):

```bash
kubectl apply -f kubernetes/blue -f kubernetes/green
kubectl port-forward svc/blue 8080:80 &
curl http://localhost:8080/
kubectl port-forward svc/green 8081:80 &
curl http://localhost:8081/
```

Or install it via Helm:

```bash
helm upgrade --install colour kubernetes/chart
```

Or just run the container directly:

```bash
docker run -p 8080:8080 -e HOSTNAME=blue-demo ghcr.io/platformfix/colour:latest
curl http://localhost:8080/
```

## How it works

The server reads two things from its environment:

- `HOSTNAME`: Kubernetes sets this to the pod's name automatically. The
  colour is everything before the first `-` (`blue-7d9f9c5db4-abcde` →
  `blue`).
- `NAMESPACE`: set explicitly, or read from the in-cluster serviceaccount
  file when running inside a real cluster.

`PORT` controls the listen port (default `8080`).

## Local development

```bash
go test ./...
golangci-lint run ./...
goreleaser build --single-target --snapshot --clean -o colour
docker build -t colour:dev .
```

## Releases

Tagged releases (`vX.Y.Z`) are built and published by
[goreleaser](.goreleaser.yaml): a multi-arch image pushed to
`ghcr.io/platformfix/colour`, cosign-signed (keyless, via GitHub's OIDC
identity), with an SBOM and SLSA build provenance attached.

## License

[MIT](LICENSE)
