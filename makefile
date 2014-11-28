help:
	@echo "build - builds artifacts into ./bin"

build:
	@cd killapp/ && \
	go build -o ../bin/killapp && \
	GOOS=windows GOARCH=386 go build -o ../bin/killapp-32.exe && \
	GOOS=windows GOARCH=amd64 go build -o ../bin/killapp-64.exe;
