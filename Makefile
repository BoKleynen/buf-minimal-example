DOCKER_REPOSITORY=example

.PHONY: dev-image
dev-image:
	@DOCKER_BUILDKIT=1 docker build \
		--target development \
		-t $(DOCKER_REPOSITORY)/dev \
		.

.PHONY: buf-generate
buf-generate: dev-image
	@docker run \
		--rm \
    	--volume "$(shell pwd):/workspace" \
    	--workdir /workspace \
    	$(DOCKER_REPOSITORY)/dev \
    	'buf generate'
