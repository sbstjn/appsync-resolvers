test:
	@ go test -cover -coverprofile=coverage.out -race ./... 

coveralls:
	@ goveralls -coverprofile=coverage.out -service=circle-ci -repotoken=$(COVERALLS_TOKEN)

.PHONY: test coverall