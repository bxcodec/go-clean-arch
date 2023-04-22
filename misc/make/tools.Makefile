# This makefile should be used to hold functions/variables

define github_url
    https://github.com/$(GITHUB)/releases/download/v$(VERSION)/$(ARCHIVE)
endef

# creates a directory bin.
bin:
	@ mkdir -p $@

# ~~~ Tools ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

# ~~ [migrate] ~~~ https://github.com/golang-migrate/migrate ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

MIGRATE := $(shell command -v migrate || echo "bin/migrate")
migrate: bin/migrate ## Install migrate (database migration)

bin/migrate: VERSION := 4.14.1
bin/migrate: GITHUB  := golang-migrate/migrate
bin/migrate: ARCHIVE := migrate.$(OSTYPE)-amd64.tar.gz
bin/migrate: bin
	@ printf "Install migrate... "
	@ curl -Ls $(call github_url) | tar -zOxf - ./migrate.$(shell echo $(OSTYPE) | tr A-Z a-z)-amd64 > $@ && chmod +x $@
	@ echo "done."

# ~~ [ air ] ~~~ https://github.com/cosmtrek/air ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

AIR := $(shell command -v air || echo "bin/air")
air: bin/air ## Installs air (go file watcher)

bin/air: VERSION := 1.15.1
bin/air: GITHUB  := cosmtrek/air
bin/air: ARCHIVE := air_$(VERSION)_$(OSTYPE)_amd64.tar.gz
bin/air: bin
	@ printf "Install air... "
	@ curl -Ls $(shell echo $(call github_url) | tr A-Z a-z) | tar -zOxf - air > $@ && chmod +x $@
	@ echo "done."


# ~~ [ gotestsum ] ~~~ https://github.com/gotestyourself/gotestsum ~~~~~~~~~~~~~~~~~~~~~~~

GOTESTSUM := $(shell command -v gotestsum || echo "bin/gotestsum")
gotestsum: bin/gotestsum ## Installs gotestsum (testing go code)

bin/gotestsum: VERSION := 1.6.1
bin/gotestsum: GITHUB  := gotestyourself/gotestsum
bin/gotestsum: ARCHIVE := gotestsum_$(VERSION)_$(OSTYPE)_amd64.tar.gz
bin/gotestsum: bin
	@ printf "Install gotestsum... "
	@ curl -Ls $(shell echo $(call github_url) | tr A-Z a-z) | tar -zOxf - gotestsum > $@ && chmod +x $@
	@ echo "done."

# ~~ [ tparse ] ~~~ https://github.com/mfridman/tparse ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

TPARSE := $(shell command -v tparse || echo "bin/tparse")
tparse: bin/tparse ## Installs tparse (testing go code)

bin/tparse: VERSION := 0.8.3
bin/tparse: GITHUB  := mfridman/tparse
bin/tparse: ARCHIVE := tparse_$(VERSION)_$(OSTYPE)_x86_64.tar.gz
bin/tparse: bin
	@ printf "Install tparse... "
	@ curl -Ls $(call github_url) | tar -zOxf - tparse > $@ && chmod +x $@
	@ echo "done."

# ~~ [ mockery ] ~~~ https://github.com/vektra/mockery ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

MOCKERY := $(shell command -v mockery || echo "bin/mockery")
mockery: bin/mockery ## Installs mockery (mocks generation)

bin/mockery: VERSION := 2.5.1
bin/mockery: GITHUB  := vektra/mockery
bin/mockery: ARCHIVE := mockery_$(VERSION)_$(OSTYPE)_x86_64.tar.gz
bin/mockery: bin
	@ printf "Install mockery... "
	@ curl -Ls $(call github_url) | tar -zOxf -  mockery > $@ && chmod +x $@
	@ echo "done."

# ~~ [ golangci-lint ] ~~~ https://github.com/golangci/golangci-lint ~~~~~~~~~~~~~~~~~~~~~

GOLANGCI := $(shell command -v golangci-lint || echo "bin/golangci-lint")
golangci-lint: bin/golangci-lint ## Installs golangci-lint (linter)

bin/golangci-lint: VERSION := 1.39.0
bin/golangci-lint: GITHUB  := golangci/golangci-lint
bin/golangci-lint: ARCHIVE := golangci-lint-$(VERSION)-$(OSTYPE)-amd64.tar.gz
bin/golangci-lint: bin
	@ printf "Install golangci-linter... "
	@ curl -Ls $(shell echo $(call github_url) | tr A-Z a-z) | tar -zOxf - $(shell printf golangci-lint-$(VERSION)-$(OSTYPE)-amd64/golangci-lint | tr A-Z a-z ) > $@ && chmod +x $@
	@ echo "done."