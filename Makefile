.PHONY: run-base run-rep-gen

run-base:
	go run cmd/crystal-golang/base-app/main.go

run-rep-gen:
	go run cmd/crystal-golang/rep-gen-app/main.go
