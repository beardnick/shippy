build:
	# include的路径
	# 在指定路径的时候可以加上其它的一些参数，比如说插件等，最后用：来分隔输出目录
	protoc -I. \
		--go_out=plugins=micro:. \
		*.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-service .

run:
	docker run -p 50051:50051 consignment-service

