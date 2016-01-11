build:
	go build .
install:
	go install .
test:
	go run ftmpl.go -targetgo example/compiled.go example
	go test -v .
