build:
	# include的路径
	# 在指定路径的时候可以加上其它的一些参数，比如说插件等，最后用：来分隔输出目录
	protoc -I. \
		--go_out=plugins=micro:$(GOPATH)/src/github.com/beardnick/shippy/consignment-service \
		proto/consignment/consignment.proto
	docker build -t consignment-service .

run:
	docker run -p 9090:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns consignment-service

