include .env
#MAKEFILE_VAR := This_is_the_way_to_create_a_makefile_var

.PHONY:

tidy: ## Runs go mod tidy
	go mod tidy

build: ## Builds Binaries on ./bin
	go build -o ./bin/main ./cmd/main.go

clean: ## Removes old builds
	@rm -rf bin/*

run: ## Runs the App
	. ./.env && go run ./cmd/main.go

docker-build: ## Create Docker Image
	@echo "Building docker image..."
	@docker build \
	-t ${PROJECT_NAME}:${VERSION} \
	.
	@docker image prune --filter LABEL=stage=auto-clean -f
	@echo "Done!"

docker-run: ## Run Docker Image as Container
	@docker run --name ${PROJECT_NAME} -d ${PROJECT_NAME}:${VERSION} -p 8080:8080

docker-down: ## Down and clean Docker Image
	@echo "Down App & Clean docker image..."
	@docker container stop ${PROJECT_NAME} 
	@docker container rm ${PROJECT_NAME}
	@docker image rm ${PROJECT_NAME}