APP_NAME=domain-deploy
APP_VERSION=1.0.0
flags=-X main.appName=$(APP_NAME) -X main.appVersion=$(APP_VERSION) -X main.main.goVersion=$(go version) -X main.main.buildTime=`date -u '+%Y-%m-%d'` -X main.gitHash=`git show -s --format=%H`

build:
	CGO_ENABLED=0 GOARCH=amd64 go build -v -trimpath -tags=jsoniter -ldflags "-s -w $(flags)" -o ./bin/${APP_NAME} ./main/main.go

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -v -trimpath -tags=jsoniter -ldflags "-s -w $(flags)" -o ./bin/${APP_NAME}-windows-amd64.exe ./main/main.go

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -trimpath -tags=jsoniter -ldflags "-s -w $(flags)" -o ./bin/${APP_NAME}-linux-amd64 ./main/main.go

darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -v -trimpath -tags=jsoniter -ldflags "-s -w $(flags)" -o ./bin/${APP_NAME}-darwin-amd64 ./main/main.go

darwin_arm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -v -trimpath -tags=jsoniter -ldflags "-s -w $(flags)" -o ./bin/${APP_NAME}-darwin-arm64 ./main/main.go