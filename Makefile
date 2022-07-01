.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help


gr-build-only: ## Go Releaser: Build (snapshot) only
	goreleaser build --snapshot --rm-dist

gr-release-snapshot-dry-run: ## Go Releaser: Release snapshot dry-run
	goreleaser release --skip-publish --rm-dist --snapshot

gr-release-dry-run: ## Go Releaser: Release dry-run (required tag)
	goreleaser release --skip-publish --rm-dist