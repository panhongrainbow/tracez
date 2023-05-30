.PHONY: help

estimate:
	go test -v -bench='^\QBenchmark_Estimate_omitemptySample' -run=none .
	@sleep 5
	go test -v -bench='^\QBenchmark_Estimate_nonOmitemptySample' -run=none .
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "The purpose of this benchmark test is primarily to"
	@echo "compare the performance between"
	@echo "nonOmitemptySample and omitemptySample."
	@echo ""
	@echo "Available targets:"
	@echo "  estimate     - benchmark test"
	@echo ""