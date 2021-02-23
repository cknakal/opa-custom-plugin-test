run-dev: ## Run the opa sidecar and envoy proxy
	@docker-compose -f docker-compose.yml up --build