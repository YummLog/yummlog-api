
.PHONY: gen-code
gen-code:
	@oapi-codegen -generate types -package model openapi.yaml > ./internal/model/model.go
	@oapi-codegen -generate gin -package model openapi.yaml > ./internal/model/server.go
	@sqlc generate