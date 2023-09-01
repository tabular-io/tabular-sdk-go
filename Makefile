build:
	openapi-generator \
		generate -i openapispec.json -g go --package-name tabular \
		--git-host github.com --git-user-id tabular-io --git-repo-id tabular-sdk-go \
		-o tabular/
	cd tabular && go mod tidy
