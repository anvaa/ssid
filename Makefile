
# GOOS=darwin GOARCH=amd64

appname=SSID
appnamelong=Super Simple Inventory Datatbase

# appname in lower case
buildname=$(shell echo $(appname) | tr '[:upper:]' '[:lower:]')
appbundlepath=../app/ssid.app/Contents/MacOS
arch=$(shell go env GOARCH)
os=$(shell go env GOOS)

build: clean

	# Build for local platform
	CGO_ENABLED=1 go build -o bin/$(buildname)_$(os).$(arch) -ldflags="-s -w -X 'app/app_conf.AppName=$(appname)' -X 'app/app_conf.AppNameLong=$(appnamelong)'" cmd/main.go

	# if file exists, copy it to the app bundle
	@if [ -f $(appbundlepath)/ssid ]; then \
		cp bin/$(buildname)_$(os).$(arch) $(appbundlepath)/ssid; \
	fi

run:
	go run cmd/main.go

runapp:
	./bin/$(buildname)_$(os).$(arch)

clean:
 	#if file exists, delete Inventory
	@if [ -f bin/$(buildname)_$(os).$(arch) ]; then \
		rm bin/$(buildname)_$(os).$(arch); \
	fi
	
