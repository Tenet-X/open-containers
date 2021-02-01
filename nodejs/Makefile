node_ver := 14.15.4
pwd := $(shell pwd)

nodejs:
	@[ -d tmp ] || mkdir tmp
	@curl -SLO "http://nodejs.org/dist/v$(node_ver)/node-v$(node_ver)-linux-x64.tar.gz" \
		&& tar -xzf "node-v$(node_ver)-linux-x64.tar.gz" -C $(pwd)/tmp/ \
		&& rm "node-v$(node_ver)-linux-x64.tar.gz" \
		&& mv $(pwd)/tmp/node-v$(node_ver)-linux-x64 $(pwd)/tmp/nodejs
	docker build -t nodejs:lts-14 -f Dockerfile.nodejs .
	@-rm -rf tmp

clean:
	@-rm -rf "node-v$(node_ver)-linux-x64.tar.gz" tmp

.PHONY: nodejs
