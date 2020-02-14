fmt:
	@echo ">> Running Gofmt.."
	gofmt -l -s -w .

get:
	@echo ">> Getting any missing dependencies.."
	go get ./...

install:
	go install github.com/nbw/palindrome

test:
	go test -v ./...





