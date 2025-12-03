SHELL := /usr/bin/env bash
.SHELLFLAGS := -Eeuo pipefail -c

BUILD_DIR   ?= build
BIN_NAME    ?= mooltipass
RELEASE_DIR ?= release
TARGET_REPO ?=        
SSH_ALIAS   ?=       
# one timestamp per `make` invocation, safe for branch names (no colons/spaces)
TIMESTAMP     := $(shell date -u +%Y%m%dT%H%M%SZ)
RELEASE_PREFIX ?= release
TARGET_BR     ?= $(RELEASE_PREFIX)-$(TIMESTAMP)

.PHONY: sanity build release release-working-tree scan push-artifacts clean

sanity:
	@# Guard against dangerous RELEASE_DIR
	@test -n "$(RELEASE_DIR)" && [ "$(RELEASE_DIR)" != "/" ] && [ "$(RELEASE_DIR)" != "." ] || { \
	  echo "ERR: unsafe RELEASE_DIR='$(RELEASE_DIR)'"; exit 2; }

build:
	@mkdir -p "$(BUILD_DIR)"
	go build -o "$(BUILD_DIR)/$(BIN_NAME)" ./cmd/mooltipass

# Preferred: export only tracked files from the index (no tar, no untracked surprises)
release: sanity
	@echo "Preparing release artifacts from HEAD (tracked files only)…"
	rm -rf -- "$(RELEASE_DIR)"
	mkdir -p "$(RELEASE_DIR)"
	git --work-tree="$(RELEASE_DIR)" checkout-index -a -f
	@echo "Release artifacts prepared in $(RELEASE_DIR)"

# Alternate: mirror working tree (includes untracked); exclude obvious junk
release-working-tree: sanity
	@echo "Preparing release artifacts from working tree (no rsync/tar)…"
	rm -rf -- "$(RELEASE_DIR)"
	mkdir -p "$(RELEASE_DIR)"
	# copy top-level entries, including dotfiles; prune sensitive dirs/files
	find . -mindepth 1 -maxdepth 1 \
	  \( -name '.git' -o -name '$(RELEASE_DIR)' -o -name '$(BUILD_DIR)' -o -name '.env' \
	     -o -name '.ssh' -o -name '*id_rsa*' -o -name '*id_ed25519*' -o -name '*.pem' -o -name 'push-to-repo' \) -prune -o \
	  -exec cp -a {} "$(RELEASE_DIR)"/ \;
	@echo "Release artifacts prepared in $(RELEASE_DIR)"

# Secret-ish quick scan; excludes .env.example and Makefile
UNAME_S := $(shell uname -s)
# treat macOS and *BSD as "BSD grep" envs; everything else → GNU grep
ifneq (,$(filter Darwin FreeBSD OpenBSD NetBSD,$(UNAME_S)))
  USE_BSD_GREP := 1
else
  USE_BSD_GREP := 0
endif

# Prefer Homebrew GNU grep on macOS if available
GGREP := $(shell command -v ggrep 2>/dev/null)

# patterns to scan for
SECRET_RE := (AWS_SECRET_ACCESS_KEY|BEGIN (RSA|EC|OPENSSH) PRIVATE KEY|password=|api[_-]?key)

.PHONY: scan
ifeq ($(USE_BSD_GREP),1)
scan:
	@if [ -n "$(GGREP)" ]; then \
	  echo "scan: using GNU grep (ggrep)"; \
	  matches="$$( $(GGREP) -R -n -I -H --exclude='.env.example' --exclude='Makefile' -E '$(SECRET_RE)' "$(RELEASE_DIR)" 2>/dev/null || true )"; \
	else \
	  echo "scan: using BSD grep"; \
	  matches="$$( find "$(RELEASE_DIR)" -type f ! -name '.env.example' ! -name 'Makefile' -print0 \
	    | xargs -0 grep -n -I -H -E '$(SECRET_RE)' 2>/dev/null || true )"; \
	fi; \
	if [ -n "$$matches" ]; then \
	  echo "$$matches"; \
	  echo "----- files with matches -----"; \
	  echo "$$matches" | awk -F: '{print $$1}' | sort -u; \
	  echo "ERR: secret-like strings found in $(RELEASE_DIR)"; exit 3; \
	else \
	  echo "scan: no matches"; \
	fi
else
scan:
	@echo "scan: using GNU grep"
	@matches="$$( grep -R -n -I -H --exclude='.env.example' --exclude='Makefile' -E '$(SECRET_RE)' "$(RELEASE_DIR)" 2>/dev/null || true )"; \
	if [ -n "$$matches" ]; then \
	  echo "$$matches"; \
	  echo "----- files with matches -----"; \
	  echo "$$matches" | awk -F: '{print $$1}' | sort -u; \
	  echo "ERR: secret-like strings found in $(RELEASE_DIR)"; exit 3; \
	else \
	  echo "scan: no matches"; \
	fi
endif


push-artifacts: release
	@echo "Publishing $(RELEASE_DIR) → $(TARGET_REPO) branch $(TARGET_BR) via alias $(SSH_ALIAS)…"
	chmod +x bin/push-to-repo
	./bin/push-to-repo "$(abspath $(RELEASE_DIR))" "$(TARGET_REPO)" "$(TARGET_BR)" "$(SSH_ALIAS)"

clean:
	rm -rf -- "$(BUILD_DIR)" "$(RELEASE_DIR)"