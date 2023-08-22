export PATH := $(GOPATH)/bin:$(PATH)
export GO111MODULE=on

MODULE = dootask-okr

PORT 			:= 5566
VERSION			:= $(shell git describe --tags --always --match="v*" 2> /dev/null || echo v0.0.0)
VERSION_HASH	:= $(shell git rev-parse --short HEAD)

GOCGO 	:= env CGO_ENABLED=1
LDFLAGS	:= -s -w -X "$(MODULE)/config.Version=$(VERSION)" -X "$(MODULE)/config.CommitSHA=$(VERSION_HASH)"

run: build
	./main --mode debug

watch:
	@if lsof -i :$(PORT) >/dev/null 2>&1; then \
        echo "Port $(PORT) is already in use, killing the process..."; \
        lsof -i :$(PORT) | awk 'NR!=1 {print $$2}' | xargs kill; \
    fi
	$(GOCGO) air

release:
	cd web && npm run build && cd ../ && rm -f main
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
	OCKER_BUILDKIT=1 docker buildx build --push -t hitosea2020/okr:0.0.3 --platform linux/amd64,linux/arm64 .

build:
	cd web && npm run build && cd ../
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go main

monitor:
	lsof -i :5567 | grep node | awk '{print $$2}' | xargs kill -9
	cd web && nohup npm run dev > ../output.log >&1 & cd ../ 
	${HOME}/go/bin/fresh -c ./fresh.conf
	
translate:
	cd web && npm run translate && cd ../

# 提示 fresh: No such file or directory 时解決辦法
# go install github.com/pilu/fresh@latest

# 提示 air: No such file or directory 时解決辦法
# go install github.com/cosmtrek/air@latest

# 提示 swag: No such file or directory 时解決辦法
# go get -u github.com/swaggo/swag/cmd/swag
# go install github.com/swaggo/swag/cmd/swag@latest
