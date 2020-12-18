# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

run-local:
	@echo '(+) Starting to build LOCAL env'
	go run main.go

setup-local:
	@echo '(+) Setting up go modules in env'
	go get -u github.com/gorilla/mux
	go get -u github.com/lib/pq
	go get -u github.com/stretchr/testify/assert
	go get -u github.com/go-sql-driver/mysql
	go get -u github.com/joho/godotenv
	@echo '(+) Modules set'
	@echo '(+) Build MySQL container golang-bootcamp-db'
	@if [ "`docker ps -a | grep golang-bootcamp-db | cut -f1 -d\ `" ]; then \
		docker start golang-bootcamp-db; \
	else \
		docker build -t golang-bootcamp-db .; \
		docker run -d -p4545:3306 --name golang-bootcamp-db golang-bootcamp-db; \
	fi
	@echo '<= MySQL container done.'
	docker exec -i golang-bootcamp-db mysql -u root -ppassword bitcoin_db < db/sample.sql


version: ## Output the current version
	@echo $(VERSION)
