swagger:
	swag init --parseDependency --generalInfo ./cmd/microddd/main.go --output ./interfaces/swagger/
run:
	go run cmd/microddd/main.go cmd/microddd/wire_gen.go
wire:
	wire cmd/microddd/*