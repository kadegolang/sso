PKG := "gitee.com/go-kade/sso"

mod: ## Get the dependencies
	@go mod tidy

runmc: ## Run Server
	@ go run main.go

install: ## Install depence go package
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/favadi/protoc-go-inject-tag@latest

gen: ## protobuf 编译
	@protoc -I=. --go_out=. --go_opt=module=${PKG} --go-grpc_out=. --go-grpc_opt=module=${PKG} */app/*/pb/*.proto
	@protoc-go-inject-tag -input=*/app/*/*.pb.go

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'