HASURAENV_TEST_CLI_PATH:=tmp/test/bin/hasuraenv

build:
	go build -o tmp/bin/hasuraenv cmd/hasuraenv/hasuraenv.go

build-for-e2e-test:
	go build \
		-o e2e/${HASURAENV_TEST_CLI_PATH} \
		-ldflags '-X main.configPathBase=tmp/test/.hasuraenv -X main.version=e2e' \
		cmd/hasuraenv/hasuraenv.go

e2e-testing:
	export HASURAENV_TEST_CLI_PATH=${HASURAENV_TEST_CLI_PATH} && \
		go test ./e2e -v -count=1
