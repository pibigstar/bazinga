all: build

PROJECT_NAME=bazinga

IMAGES_VERSION=0.0.1

IMAGES_NAME=github.com/pibigstar/bazinga:$(IMAGES_VERSION)

CODE_PATH=internal/code

PROTO_FILE=app/grpc/proto/bazinga.proto

clean:
	@rm -rf bin/$(PROJECT_NAME)
	@rm -rf log/$(PROJECT_NAME)

build:
	@go build -o bin/$(PROJECT_NAME) main.go
	@chmod +x bin/$(PROJECT_NAME)

build-docker-images:
	@docker build -t $(IMAGES_NAME) .

push-docker-images:
	@docker push $(IMAGES_NAME)

docker-deploy:
	@if [[ -d "docker ps -a | grep $(PROJECT_NAME)" ]]; then docker rm -f grep $(PROJECT_NAME) ; fi
	@docker run --name $(PROJECT_NAME) -p 8080:80 $(IMAGES_NAME)

# 生成code码
code:
	@protoc --go_out=plugins=grpc,paths=import:./ $(CODE_PATH)/code.proto
	@echo "// Code impl\nfunc (e Error) Code() int {\n	return int(e) \n}" >> $(CODE_PATH)/code.pb.go

# 生成protoc
proto:
	@protoc --go_out=plugins=grpc,paths=import:./ $(PROTO_FILE)