cov:
	go test -count=1 . -coverprofile=cov.out
	go tool cover -html=cov.out

test:
	go test -v -count=1 .
