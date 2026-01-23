export GOPROXY=https://goproxy.cn/



default:
	@rm -rf dist && mkdir dist
	@rm /Users/tongxu/go/bin/linux_amd64/shop*
	@make build
	@cp -rf /Users/tongxu/go/bin/linux_amd64/shop* ./dist
	@rm -rf  /Users/tongxu/Projects/meepod/opt/projects/shop/dist/backend/ && mkdir -p /Users/tongxu/Projects/meepod/opt/projects/shop/dist/backend/
	@cp -rf /Users/tongxu/go/bin/linux_amd64/shop* /Users/tongxu/Projects/meepod/opt/projects/shop/dist/backend/

release:
	@make docker-build
	@make docker-push

build:
	@go mod tidy

	@go vet ./...
	@echo "static analysis of Go program ... done. "

	@go test ./...
	@echo "testing of Go program ... done. "

	@GOOS=linux GOARCH=amd64 go build -o ../../opt/projects/shop/dist/backend/ ./cmd/...
	@echo "building of Go program ... done."

docker-build:
	@docker build -t hub.lottery.local/shop/shop_rest:${version} --build-arg SERVER_NAME=shop_rest .
	@docker build -t hub.lottery.local/shop/shop_worker:${version} --build-arg SERVER_NAME=shop_worker .
	@docker build -t hub.lottery.local/shop/shop_pay_rest:${version} --build-arg SERVER_NAME=shop_pay_rest .
	@docker build -t hub.lottery.local/shop/shop_pay_worker:${version} --build-arg SERVER_NAME=shop_pay_worker .
	@docker build -t hub.lottery.local/shop/shop_admin_rest:${version} --build-arg SERVER_NAME=shop_admin_rest .


docker-push:
	@docker login https://hub.lottery.local -u admin -p bitnami
	@docker push hub.lottery.local/shop/shop_rest:${version}
	@docker push hub.lottery.local/shop/shop_worker:${version}
	@docker push hub.lottery.local/shop/shop_pay_rest:${version}
	@docker push hub.lottery.local/shop/shop_pay_worker:${version}
	@docker push hub.lottery.local/shop/shop_admin_rest:${version}
