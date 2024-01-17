test-coverage:
	@echo "=================================================================================="
	@echo "Coverage Test"
	@echo "=================================================================================="
	go test -v -coverprofile coverage.cov ./...
	@echo "\n"
	@echo "=================================================================================="
	@echo "All Package Coverage"
	@echo "=================================================================================="
	go tool cover -func coverage.cov

build:
	@echo "=================================================================================="
	@echo "Build Service"
	@echo "=================================================================================="
	docker-compose up

delete:
	@echo "=================================================================================="
	@echo "Stop Service"
	@echo "=================================================================================="
	docker stop chatbot-golang-app-1 || true
	docker stop chatbot-mysql-1 || true
	docker stop chatbot-redis-1 || true
	
	@echo "=================================================================================="
	@echo "Delete Container"
	@echo "=================================================================================="
	docker rm chatbot-golang-app-1 || true
	docker rm chatbot-mysql-1 || true
	docker rm chatbot-redis-1 || true
	
	@echo "=================================================================================="
	@echo "Delete Images"
	@echo "=================================================================================="
	docker rmi chatbot-golang-app || true
	docker rmi redis || true
	docker rmi mysql || true