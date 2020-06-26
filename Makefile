PHONY: .compile
compile: build/_output/bin/kubeserial

build/_output/bin/kubeserial:
	@mkdir -p build/_output/x86bin/
	@go build -o build/_output/x86bin/kubeserial cmd/manager/main.go

PHONY: .build
clean:
	@rm -rf build/_output