build:
	GO_POST_PROCESS_FILE="gofmt -w" openapi-generator \
		generate -i openapispec.json -g go --package-name tabular  \
		--git-host github.com --git-user-id tabular-io --git-repo-id tabular-sdk-go \
		--additional-properties=isGoSubmodule=true \
		-o tabular/
