PROTOC ?= protoc

SRCDIR ?= proto
OUTDIR ?= $(shell cd ../../.. && pwd)

SOURCE = $(SRCDIR)/remote_cdm_session_data.proto

all: $(SOURCE)
	$(PROTOC) --proto_path=$(SRCDIR) --go_out=$(OUTDIR) $(SOURCE)
