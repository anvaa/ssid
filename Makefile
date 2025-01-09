
# GOOS=darwin GOARCH=amd64

appname=SSID
appnamelong=Super Simple Inventory Datatbase

buildname=ssid
arch=$(shell go env GOARCH)
os=$(shell go env GOOS)

build: clean

	# Build for local platform
	CGO_ENABLED=1 go build -o bin/$(buildname)_$(os).$(arch) -ldflags="-s -w -X 'app/app_conf.AppName=$(appname)' -X 'app/app_conf.AppNameLong=$(appnamelong)'" cmd/main.go

run:
	go run cmd/main.go

runapp:
	./bin/$(buildname)_$(os).$(arch)

clean:
	rm -rf bin/*