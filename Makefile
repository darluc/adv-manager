GO=go
CP=cp

.PHONY: clean server

server:./server/*.go
	$(GO) build -o ./dist/adv-manager ./server/*.go
	$(CP) config.json ./dist/config.json
	$(CP) -r public ./dist/
clean:
	-rm -rf ./dist/*