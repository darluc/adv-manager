GO=go
CP=cp

.PHONY: clean server docker

docker:server
	docker build -t 'darluc/adv-manager:1.0' -f docker/Dockerfile .
server:./server/*.go
	$(GO) build -o ./dist/adv-manager ./server/*.go
	$(CP) config.json ./dist/config.json
	$(CP) -r public ./dist/
clean:
	-rm -rf ./dist/*