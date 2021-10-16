MOCKS_DESTINATION=mocks
.PHONY: mocks
mocks: $(shell find ./pkg -name "*.go")
	@echo "Generating mocks..."
	@rm -rf $(MOCKS_DESTINATION)
	@for file in $^; do mockgen -source=$$file -destination=$(MOCKS_DESTINATION)/$$file; done