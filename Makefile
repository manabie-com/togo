
## Generate Dependency
wire:
	@cd internal/providers && wire

## Generate Swagger Document
doc:
	@swag init