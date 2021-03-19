APP_NAME = ArabicToRomanConverter

.PHONY: test
test:
	@echo '->' testing has been started ...
	go test ./service -v

.PHONY: bench
bench:
	@echo ''
	@echo '->' benchmark has been started ...
	go test -bench= ./service -v

.PHONY: run
run:
	@echo ''
	@echo '->' The application '"'$(APP_NAME)'"' has been started ...
	go run main.go

.PHONY: explore
explore: test bench
