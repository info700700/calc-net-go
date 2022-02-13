BINARY_NAME=calc-server

build:
	go build -o ${BINARY_NAME}

build_image:
	docker build --tag calc-server .
