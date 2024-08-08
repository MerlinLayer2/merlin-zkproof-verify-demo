maDEP := $(shell command -v dep 2> /dev/null)
SUM := $(shell which shasum)

export GO111MODULE=on

build_tags = netgo

ifeq ($(WITH_CLEVELDB),yes)
  build_tags += gcc
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

BUILD_FLAGS := -ldflags '$(ldflags)'

build:
	go build $(BUILD_FLAGS) -o verify .


.PHONY: build
