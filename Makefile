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
	docker stop chatbot-app || true
	docker stop mysql-app || true
	docker stop redis-app || true
	
	@echo "=================================================================================="
	@echo "Delete Container"
	@echo "=================================================================================="
	docker rm chatbot-app || true
	docker rm mysql-app || true
	docker rm redis-app || true
	
	@echo "=================================================================================="
	@echo "Delete Images"
	@echo "=================================================================================="
	docker rmi chatbot_golang-app:latest || true
	docker rmi golang:latest || true
	docker rmi redis:latest || true
	docker rmi mysql:latest || true