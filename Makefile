build:
	go build -o calc-server

clean:
	rm ./calc-server

build_image:
	docker build --tag calc-server .
