build: 
	@go build -o plugins/hello plugin/hello_impl.go
	@go build -o plugins/plugin plugin/plugin_impl.go

run:
	@go run main.go

deps:
	@go get -u \
		github.com/golang/dep/cmd/dep
	@dep ensure
