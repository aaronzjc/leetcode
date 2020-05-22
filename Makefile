GO111MODULE=on

test:
	go fmt ./...
	go test -cover ./...

coverage.out:
	go test -coverprofile=coverage.out ./...

test_profile:clean coverage.out
	go tool cover -func=coverage.out

test_profile_html:clean coverage.out
	go tool cover -html=coverage.out

.PHONY: clean
clean:
	-rm coverage.out