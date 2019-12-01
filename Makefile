TAG := $$( git rev-parse --short HEAD )
IMAGE := lms-backend\:${TAG}

migrate_up=go run main.go migrate --direction=up --step=0
migrate_down=go run main.go migrate --direction=down --step=0


.PHONY: proto run-dev build push-image build-image run-image

migrate:
	@if [ "$(DIRECTION)" = "" ] || [ "$(STEP)" = "" ]; then\
		$(migrate_up);\
	else\
		go run main.go migrate --direction=$(DIRECTION) --step=$(STEP);\
	fi

serve:
	go run . server

build:
	@GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o bin/app .

build-image:
	@make build
	@docker build -f ./Dockerfile -t ${IMAGE} .

run-image:
	@docker run --name article -p 9000:9000 ${IMAGE}

