GO=go
CP=cp

.PHONY: clean server docker

docker:
	CGO_ENABLED=1 $(GO) build -a -installsuffix cgo -ldflags '-s' -o ./dist/adv-manager ./server/*.go
	docker build -t 'darluc/adv-manager:1.1' -f docker/Dockerfile .


server:
	$(GO) build -o ./dist/adv-manager-server ./server/*.go
	$(CP) -r public ./dist/
clean:
	-rm -rf ./dist/*