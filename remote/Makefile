PROTOC ?= protoc

SRCDIR ?= proto
OUTDIR ?= $(shell cd .. && pwd)

MODULE = github.com/go-webdl/drm

SOURCE = $(SRCDIR)/remote_cdm_session_data.proto

all: $(SOURCE)
	$(PROTOC) --proto_path=$(SRCDIR) --go_out=$(OUTDIR) --go_opt=module=$(MODULE) $(SOURCE)
