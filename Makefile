all:
	bower install
	go generate ./...
	go install github.com/andrei-m/bank/bankcmd
