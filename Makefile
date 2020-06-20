build:
	cd api &&set GOOS=linux&&set GOARCH=amd64&&go build -o micro
	cd service/greeter && protoc --proto_path=$(GOPATH)/src:proto --micro_out=proto --go_out=proto greeter.proto &&set GOOS=linux&&set GOARCH=amd64&&go build
	cd service/user && protoc --proto_path=$(GOPATH)/src:proto --micro_out=proto --go_out=proto user.proto &&set GOOS=linux&&set GOARCH=amd64&&go build
	cd config && ./gradlew build
	docker-compose build

run:
	docker-compose up --force-recreate
