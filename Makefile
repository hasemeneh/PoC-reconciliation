pkgs          = $(shell go list ./... | grep -vE '(vendor|mock)')
NOW=$(shell date)
compose_file=./docker/docker-compose.yml
compose=docker-compose -f ${compose_file}
reconcile_service_binary=reconcile
current_dir=$(shell pwd)


docker-start:
	@echo "${NOW} STARTING CONTAINER..."
	@${compose} up -d --build

docker-rebuilddb-reconcile:
	@echo "${NOW} REBUILDDB..."
	@echo "${NOW} DROPING EXISTING DB..."
	@docker exec -it reconcile-db  mysql -uroot -e'drop database if exists reconcile_db'
	@echo "${NOW} CREATE DB..."
	@docker exec -it reconcile-db  mysql -uroot -e'create database reconcile_db'
	@echo "${NOW} RUN SQL SCRIPTS..."
	@docker exec -it reconcile-db setup.sh /etc/database

docker-stop:
	@echo "${NOW} STOPPING CONTAINER..."
	@${compose} down
	@echo "${NOW} CLEAN UP..."
	@rm -f ./bin/reconcile/${reconcile_service_binary} 

docker-run-reconcile:
	@echo "${NOW} BUILDING..."
	@cd ./svc/reconcileapp/src && go mod tidy && go build -o ./../../../bin/reconcile/${reconcile_service_binary}
	@echo "${NOW} RUNNING..."
	@docker exec -it reconcileapp /usr/local/bin/${reconcile_service_binary}

generate-mock:
	@echo "${NOW} GENERATING MOCKS..."
	@cd ./svc/reconcileapp/src && mockgen -source=./repositories/reconcilement.go -destination=./mock/mockrepository/reconcilement.go -package=mockrepository
	@cd ./svc/reconcileapp/src && mockgen -source=./repositories/saldos.go -destination=./mock/mockrepository/saldos.go -package=mockrepository
	@cd ./svc/reconcileapp/src && mockgen -source=./repositories/transactions.go -destination=./mock/mockrepository/transactions.go -package=mockrepository
	@cd ./svc/reconcileapp/src && mockgen -source=./repositories/user.go -destination=./mock/mockrepository/user.go -package=mockrepository
	
docker-exec:
	@echo "${NOW} EXECUTING COMMAND..."
	@docker exec -it reconcileapp /bin/sh -c bash