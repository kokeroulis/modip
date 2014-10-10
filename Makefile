start:
	@./scripts/helper.sh run
.PHONY: install

install:
	@./scripts/helper.sh install
.PHONY: install

install-go:
	@./scripts/helper.sh install-go
.PHONY: install-go

install-js:
	@./scripts/helper.sh install-js
.PHONY: install-js

db-create-development:
	@./scripts/helper.sh db-create development
	@./scripts/helper.sh testdata
.PHONY: db-create-development

db-create:
	@./scripts/helper.sh db-create
.PHONY: db-create

fmt:
	@gofmt -d=true -w=true .
.PHONY: fmt

vet:
	@go vet ./...
.PHONY: vet

