build:
	docker compose build kafka-api kafka-consumer
run:
	docker compose up kafka-api kafka-consumer
push:
	docker tag kafka-api:1.0.0 ghcr.io/chatchawansama/kafka-api:1.0.0
	docker push ghcr.io/chatchawansama/kafka-api:1.0.0
	docker tag kafka-consumer:1.0.0 ghcr.io/chatchawansama/kafka-consumer:1.0.0
	docker push ghcr.io/chatchawansama/kafka-consumer:1.0.0