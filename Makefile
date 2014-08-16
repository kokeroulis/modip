start:
	@./scripts/helper.sh run
.PHONY: install

install:
	@./scripts/helper.sh install
.PHONY: install

build-tools:
	@./scripts/helper.sh build-tools
.PHONY: build-tools

test:
	@./scripts/helper.sh test
.PHONY: test

db-create-development:
	@./scripts/db.sh development
	@./scripts/helper.sh testdata
.PHONY: db-create-development

db-create:
	@./scripts/db.sh production
.PHONY: db-create

