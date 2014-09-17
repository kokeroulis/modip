start:
	@./scripts/helper.sh run
.PHONY: install

install:
	@./scripts/helper.sh install
.PHONY: install

test:
	@./scripts/helper.sh test
.PHONY: test

test-web:
	@./scripts/helper.sh test-web
.PHONY: test-web

db-create-development:
	@./scripts/db.sh development
	@./scripts/helper.sh testdata
.PHONY: db-create-development

db-create:
	@./scripts/db.sh production
.PHONY: db-create

fmt:
	@gofmt -d=true -w=true .
.PHONY: fmt

vet:
	@go vet ./...
.PHONY: vet
