.PHONY: run, seed

run:
	go run cmd/crystal-golang/main.go

seed:
	go run utils/test_seed.go
