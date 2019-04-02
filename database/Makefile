CONTAINER_NAME := "doraku-db"

all: run

run:
	docker run -p 27017:27017 --name $(CONTAINER_NAME) -d mongo
	docker cp hobby.csv $(CONTAINER_NAME):hobby.csv
	docker exec -it $(CONTAINER_NAME) mongoimport --db doraku --collection hobby --type csv --file hobby.csv --headerline

rm:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)