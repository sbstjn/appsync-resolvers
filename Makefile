test:
	@ go test -cover -coverprofile=c.out -race ./... 

.PHONY: test