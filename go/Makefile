pwd := $(shell pwd)
ipaddr := $(shell hostname -I | cut -d ' ' -f 1)

# ============ Golang builder ==============
# Golang env builder
gobuilder: $(pwd)/Dockerfile.d/Dockerfile.builder.go
	@docker build \
		-t gobuilder \
		-f $(pwd)/Dockerfile.d/Dockerfile.builder.go .

gobuilder_rebuild: $(pwd)/Dockerfile.d/Dockerfile.builder.go
	@docker build --no-cache \
		-t gobuilder \
		-f $(pwd)/Dockerfile.d/Dockerfile.builder.go .

# ============ Binance Smart Chain ==============
# BSC builder
bsc_builder: gobuilder $(pwd)/Dockerfile.d/Dockerfile.builder.bsc
	@docker build \
		-t bsc_builder \
		-f $(pwd)/Dockerfile.d/Dockerfile.builder.bsc .

# build the BSC binary
build_bsc: bsc_builder
	@[ -d $(pwd)/builder_bsc ] || mkdir $(pwd)/builder_bsc
	cp $(pwd)/bin/bsc_build.sh $(pwd)/builder_bsc/bsc_build.sh \
		&& chmod +x $(pwd)/builder_bsc/bsc_build.sh
	@docker run --name bsc_builder \
		-v $(pwd)/builder_bsc:/builder:rw \
		bsc_builder /builder/bsc_build.sh
	@docker rm bsc_builder

# build the BSC container/runner
bsc_runner: build_bsc $(pwd)/Dockerfile.d/Dockerfile.runner.bsc
	@docker build \
		-t bsc_runner\
		-f $(pwd)/Dockerfile.d/Dockerfile.runner.bsc .

# clean the build, use with cautious
clean:
	@-sudo rm -rf builder_*

.PHONY: clean
