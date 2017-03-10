docs: test
	embedmd -w README.md

test:
	go test -v
