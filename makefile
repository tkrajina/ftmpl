build:
	go build .
install:
	go install .
test:
	go run ftmpl.go -targetgo example/compiled.go example
	go test -v .

install-tools:
	go get -u github.com/fzipp/gocyclo
	go get -u github.com/golang/lint
	go get -u github.com/kisielk/errcheck
	go get -u github.com/tkrajina/ftmpl
	go get -u github.com/tkrajina/gojson2models

vet:
	go tool vet --all -shadow=true . 2>&1 | grep -v "declaration of err shadows"
lint:
	golint ./... | grep -v "or be unexported" | grep -v "underscores in"
errcheck:
	errcheck ./...
gocyclo:
	-gocyclo -over 10 .
check: test gocyclo vet lint errcheck
	echo "OK"
