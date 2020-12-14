# Test database as docker container

## How to build
```
docker build -t test-db .
```
 
## How to run
```
docker run -d --name test_db -p 5432:5432 \
-e POSTGRES_USER=dev \
-e POSTGRES_PASSWORD=dev114 \
-e POSTGRES_DB=devdb \
test-db
```