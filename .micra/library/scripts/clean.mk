
.PHONY: clean
clean: $(CLEANERS)
	@echo "Cleaning intermediate build files"
	@$(GO) clean -r -cache -testcache ./...
	@find . -type f -name '*.gen.go' -exec rm {} \; 2>/dev/null

.PHONY: help-clean
help-clean:
	@echo " clean               clean intermediary files"
