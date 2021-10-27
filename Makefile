generate-client:
	find . -not -path '*/\.*' -type f \
		! -name Makefile \
		! -name LICENSE \
		! -name go.mod \
		! -name go.sum \
		-exec rm -rf {} \; \
	&& find . -depth ! -path '*/\.*' ! -path . -type d -exec rm -rf {} \; \
	&& swagger generate client -A one-paas \
		--skip-models --existing-models=github.com/onepaas/onepaas/pkg/api/v1 \
		-f ../onepaas/api/swagger.json \
		-t ./ --template=stratoscale \
		--client-package=onepaas \
	&& mv onepaas/* . \
	&& rm -rf onepaas \
	&& sed 's/onepaas.go\/onepaas/onepaas.go/g' -i one_paas_client.go \
	&& go get -u ./... \
	&& go mod tidy