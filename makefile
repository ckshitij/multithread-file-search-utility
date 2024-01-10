build:
	go build -o mfs && chmod +x mfs
run: build
	./mfs --directory <dirname> --filename <filename>