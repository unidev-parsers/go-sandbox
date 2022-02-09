GO_FILES?=$$(find . -name '*.go' | grep -v vendor)

fmt:
	@echo "Formatting files"
	gofmt -w $(GO_FILES)
	goimports -w $(GO_FILES)
