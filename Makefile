.PHONY: build
build:
	go build -o calc-server


.PHONY: clean
clean:
	rm ./calc-server


.PHONY: build_image
build_image:
	docker build --tag calc-server .
