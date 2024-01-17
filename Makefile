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

delete-service:
	@echo "=================================================================================="
	@echo "Stop Service"
	@echo "=================================================================================="
	docker stop chatbot_mysql_1
	docker stop chatbot_redis_1
	docker stop chatbot_golang-app_1
	
	@echo "=================================================================================="
	@echo "Delete Container"
	@echo "=================================================================================="
	docker rm chatbot_mysql_1
	docker rm chatbot_redis_1
	docker rm chatbot_golang-app_1
	
	@echo "=================================================================================="
	@echo "Delete Images"
	@echo "=================================================================================="
	docker rmi chatbot_golang-app
	docker rmi redis
	docker rmi mysql