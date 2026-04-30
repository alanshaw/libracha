gen:
	cd ./capabilities/assert/datamodel/gen && go run ./main.go
	cd ./capabilities/blob/datamodel/gen && go run ./main.go
	cd ./capabilities/claim/datamodel/gen && go run ./main.go
	cd ./capabilities/content/datamodel/gen && go run ./main.go
	cd ./capabilities/datamodel/gen && go run ./main.go
	cd ./capabilities/debug/datamodel/gen && go run ./main.go
	cd ./capabilities/ucan/datamodel/gen && go run ./main.go
	cd ./capabilities/upload/datamodel/gen && go run ./main.go

.PHONY: gen