# Genrator

`genrator` is a tool to help developer to generate all dockerfiles, docker-compose and jenkinsfile for every environment (development, staging and production) which fit with GCP Kuberneter Deployment.

## Installation
The best way to install this tool is downloading the binary file from newest release
__Linux and macOS:__


Please refer to our relase tag page, and download specific executable release based on platform you use.
[Release](https://gitlab.tanihub.com/tanigroup/genrator/tags)

## How to use
```sh
./genrator

What is your project name (GCP Project name) ? tanigroup-postcredit
What is your base image name (Your base image) ? tanigroup/alpine-nodejs:latest
What is your image name (Your image and container name) ? marksman
What is your exposed container port ? 3000
What is your host to container port ? 3000
Initializing Project....
Creating Dockerfiles....
Creating Compose Files....
Creating Jenkinsfiles....
Done....

```