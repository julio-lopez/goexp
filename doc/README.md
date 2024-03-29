# Notes

## CI Setup

Using GitHub Actions
* Add a YAML workflow in the [`.github/workflows`](../.github/workflows/) directory.
* Set environment => Go, version, Docker, etc. in workflow file.

## Building Container Images

There are many ways of doing this: goreleaser, simple scripts, etc.

The approach explored here uses simple shell scripts.

### Using Shell Scripts

#### Building Hello Container Image

* `build-hello.sh` compiles and builds the executuble from Go sources
* The [docker-build-hello.sh](../scripts/docker-build-hello.sh) script builds
  the container using the [`Dockerfile`](../cmd/hello/Dockerfile) in the `cmd/hello` directory. The `Dockerfile`
  simply copies the executable created by the previous script. The base image
  is Alpine
* `docker-push-hello.sh` pushes the image to Dockerhub.

The scripts are called from different stages during the GHA workflows. See the
`idle-image.yaml` for an example.
