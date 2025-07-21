#SORT_DIR := ./sort
SORT_DIR := ./list

test:
	go test $(SORT_DIR)

test-cover:
	go test -cover $(SORT_DIR)

benchmark:
	go test -bench=. $(SORT_DIR)

cover-report:
	go test -coverprofile=coverage.out $(SORT_DIR)

view-cover: cover-report
	go tool cover -html=coverage.out

clean:
	rm -f coverage.out

help:
	@echo "Available command options:"
	@echo "  make test           - Run normal tests"
	@echo "  make test-cover     - Run normal tests and display coverage"
	@echo "  make benchmark      - Run benchmarks"
	@echo "  make cover-report   - Generate a coverage report file"
	@echo "  make view-cover     - Visualize the coverage report"
	@echo "  make clean          - Clean up the generated coverage report file"
	@echo "  make help           - Display this help message"

.PHONY: test test-cover benchmark cover-report view-cover clean help