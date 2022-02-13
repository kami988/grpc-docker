go-setup:
	go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest &
	go install github.com/ramya-rao-a/go-outline@latest &
	go install github.com/cweill/gotests/gotests@latest &
	go install github.com/fatih/gomodifytags@latest &
	go install github.com/josharian/impl@latest &
	go install github.com/go-delve/delve/cmd/dlv@latest &
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest &
	go install golang.org/x/tools/gopls@latest &
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26 &
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1 &
	go install github.com/haya14busa/goplay/cmd/goplay@latest &
	go install honnef.co/go/tools/cmd/staticcheck@latest &
	go install golang.org/x/tools/cmd/goimports@latest

PROTOC_ZIP := protoc-3.15.8-linux-x86_64.zip
install-grpc:
	apt-get update
	apt-get install -y sudo
	sudo apt-get install zip unzip
	curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.15.8/$(PROTOC_ZIP)
	sudo unzip -o $(PROTOC_ZIP) -d /usr/local bin/protoc
	sudo unzip -o $(PROTOC_ZIP) -d /usr/local 'include/*'
	rm -f $(PROTOC_ZIP)

setup:
	make go-setup &
	make install-grpc

generate:
	protoc --go_out=. --go_opt=paths=source_relative ./1.18/11-workspace/def/*.proto
