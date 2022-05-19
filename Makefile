NAME=kafka-bench-test

set-kafka:
	docker compose up -d

add-topic:
	docker container exec -it broker kafka-topics --bootstrap-server broker:9092 --create --topic benchtest

clean-kafka:
	docker compose down --volumes

build:
	docker image build -t ${NAME}:dev --platform=amd64 .

build-main:
	docker image build -t ${NAME}:main --platform=amd64 -f main.Dockerfile .

run-main:
	docker container run -it --rm --name ${NAME}-main --net=host ${NAME}:main

test:
	docker container run -it --rm --name ${NAME} -v ${shell pwd}:/home/app --net=host ${NAME}:dev

.PHONY: set-kafka clean-kafka build test