# Notes

## Travis Setup

* Add [`.travis.yml`](../.travis.yml)
* Set environment => Go, version, Docker, etc.

## Docker Builds

There are many ways of doing this: goreleaser, etc. Approaches explored here

### Simple script based approach

#### Building Hello Container Image

* `build-hello.sh` compiles and builds the executuble from Go sources
* The [docker-build-hello.sh](../scripts/docker-build-hello.sh) script builds
  the container using the Dockerfile in the `hello` directory. The `Dockerfile`
  simply copies the executable created by the previous script. The base image
  is Alpine
* `docker-push-hello.sh` pushes the image to Dockerhub

The scripts are called from different stages during Travis CI. See the
`travis.yml` for details

