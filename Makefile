swagger:
	swag init --parseDependency --generalInfo ./cmd/microddd/main.go --output ./interfaces/swagger/
run:
	clear 
	go run cmd/microddd/main.go cmd/microddd/wire_gen.go
wire:
	# rm  cmd/microddd/wire_gen.go
	wire cmd/microddd/*