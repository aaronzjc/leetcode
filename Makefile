GO111MODULE=on

test:
	go fmt ./...
	go test -cover -coverprofile=coverage.out ./...

test_profile:clean test
	go tool cover -func=coverage.out

test_profile_html:clean test
	go tool cover -html=coverage.out

.PHONY: clean
clean:
	-rm coverage.out