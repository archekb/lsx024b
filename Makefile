#!/usr/bin/make
# Makefile readme (ru): <http://linux.yaroslavl.ru/docs/prog/gnu_make_3-79_russian_manual.html>
# Makefile readme (en): <https://www.gnu.org/software/make/manual/html_node/index.html#SEC_Contents>

SHELL = /bin/sh
LDFLAGS = "-s -w -X main.version=$(shell git rev-parse HEAD)"
APP_NAME = $(notdir $(CURDIR))

.PHONY : help \
         build build_arm32 \
         qemu \
         clean
.DEFAULT_GOAL : help
.SILENT : gotest

help: ## Show this help
	@printf "\033[33m%s:\033[0m\n" 'Available commands'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "  \033[32m%-11s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

qemu: ## Init qemu emulation for build application for platforms unless amd64
	docker run --privileged --rm tonistiigi/binfmt --install all

build: ## Build default version of current platform
	docker build -t $(APP_NAME):default .
	docker run --rm -v $$(pwd):/app $(APP_NAME):default go build -ldflags=$(LDFLAGS) -o $(APP_NAME) .

build_arm32v7: ## Create builder image and Build arm32v7 version of application (use for fast develop)
	docker buildx build -t $(APP_NAME):linux_arm32v7 --platform linux/arm/v7 .
	docker run --rm -v $$(pwd):/app --platform linux/arm/7 $(APP_NAME):linux_arm32v7 go build -ldflags=$(LDFLAGS) -o $(APP_NAME) .

build_web_client: ## Build web client
	docker run --rm -v "$$(pwd)/web_src:/app/web_src:ro" -v "$$(pwd)/web:/app/dist" --workdir /app --entrypoint /bin/bash node:16 -c "cp -r web_src/* . && yarn install && yarn build"

clean: ## Make clean cache
	-docker rmi $(APP_NAME):default $(APP_NAME):linux_arm32v7 node:16 -f