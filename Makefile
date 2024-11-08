APP_NAME=swift-school
GO_LINT=golangci/golangci-lint
TAG=0.0.1

.PHONY: lint
# lit: lrun lineter and static analysis tools to detect potential errors, bugs, and style issues in the project
lint:
	@echo "Linting..."
	go mod vendor
	docker run --rm -v $(shell pwd):/app -w /app ${GO_LINT} golangci-lint run --fix -v


.PHONY update-dependencies
# update-dependencies: updates all dependencies in the project
update-dependencies:
	go get -u all

.PHONY docker-run
# docker-run: runs the project in a docker container
docker-run:
	docker-compose -f tests/e2e/docker-compose.yml up --build

.PHONY docker-stop
# docker-stop: stops the project in a docker container
docker-stop:
	docker-compose -f tests/e2e/docker-compose.yml down

.PHONY test
# test: runs the tests in the project
test:
	go test ./... -cover -v


.PHONY docker-build
# docker-build: builds the project in a docker container
docker-build:
	docker build -t app .

.PHONY trivy
# trivy: scans the project for vulnerabilities
trivy:
	@if command -v some_command >/dev/null 2>&1; then\
		curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin v0.57.0;\
	else \
		echo "trivy is installed"; \
	fi

.PHONY scan
# scan: scans the project for vulnerabilities
scan:
	@echo "Scanning..."
	trivy d-build
	trivy fs.
	trivy image app

# help: print this help
help:
	@echo
	@echo " Choose a command to run in "${APP_NAME}":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo