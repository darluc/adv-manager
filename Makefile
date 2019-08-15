GO=go
CP=cp

.PHONY: clean server docker

docker:
	CGO_ENABLED=0 $(GO) build -a -installsuffix cgo -ldflags '-s' -o ./dist/adv-manager ./server/*.go
	docker build -t 'darluc/adv-manager:1.0' -f docker/Dockerfile .


server:./server/*.go
	$(GO) build -o ./dist/adv-manager-server ./server/*.go
	$(CP) -r public ./dist/
clean:
	-rm -rf ./dist/*