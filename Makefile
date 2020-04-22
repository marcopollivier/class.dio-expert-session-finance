build-image:
	docker build -t marcopollivier/finance .

push-image:
	docker push marcopollivier/finance

run-app:
	docker-compose -f .devops/app.yml up -d

stop-app:
	docker-compose -f .devops/app.yml down

prepare-env:
	docker-compose -f .devops/postgres.yml up -d

test:
	go test ./...
	go vet ./...

lint:
	golint ./...
