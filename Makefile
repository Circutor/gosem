#
# Copyright (c) 2021 Circutor S.A.
#

lint:
	golangci-lint run

test:
	go test -coverprofile=profile.cov ./...
	go tool cover -func profile.cov
	rm profile.cov
	go vet ./...
	gofmt -l .

