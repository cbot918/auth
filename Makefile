# PORT
DEVPORT = 8181
MYSQL_PORT = 8182


# dev
DEVPATH = auth.go
dev: $(DEVPATH)
	gin --appPort $(DEVPORT) run $(DEVPATH)

run: $(DEVPATH)
	go build .
	./auth.exe

## database
# instance
MYSQL_NAME = cbot-mysql
MYSQL_PASS = 12345

mysql-init:
	docker run --name $(MYSQL_NAME) -e MYSQL_ROOT_PASSWORD=$(MYSQL_PASS) -d -p $(MYSQL_PORT):3306 mysql

mysql-start:
	docker container start $(MYSQL_NAME)

mysql-client:
	docker exec -it $(MYSQL_NAME) bash

mysql-remove:
	docker stop cbot-mysql
	docker container rm cbot-mysql

# migrate
db-migrate:
	mkdir -p db/migration
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "mysql://root:12345@localhost:8201/practice?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

	
	

# apitest
APITESTENVNAME = pytest
APITESTNAME = authtest
APITESTENVWORKDIR = /app/apitest
APITESTHOSTNAME = auth.api.test

apitest-image: ./test/apitest/Dockerfile
	docker build -t $(APITESTENVNAME) ./test/apitest/.

apitest-container: 
	docker run -it -d --name $(APITESTNAME) --add-host=$(APITESTHOSTNAME):host-gateway $(APITESTENVNAME)

apitest-run: ./test/apitest/testcase
	docker exec $(APITESTNAME) rm -rf testcase
	docker exec $(APITESTNAME) mkdir testcase
	docker cp ./test/apitest/testcase/test_apitest.py $(APITESTNAME):$(APITESTENVWORKDIR)/testcase
	docker exec $(APITESTNAME) pytest --capture=no


# unittest
