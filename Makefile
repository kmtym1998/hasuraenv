HASURAENV_TEST_CLI_PATH:=tmp/bin/hasuraenv

build-for-e2e-test:
	go build -o ${HASURAENV_TEST_CLI_PATH} cmd/hasuraenv/hasuraenv.go

e2e-testing:
	export HASURAENV_TEST_CLI_PATH=${HASURAENV_TEST_CLI_PATH} && \
		go test ./e2e -v -count=1
