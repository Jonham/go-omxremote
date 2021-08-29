export GOARCH=arm
export GOOS=linux

all:PiExe

assets:
	@esc -o clientAssets.go -prefix="dist/app" dist
	@echo "clientAssets.go updated"

omx:assets
	go build -o temp/OMXRemote .

prepare:assets
	@echo "now you can run 「go run .」"

PiExe:prepare
	go build -o temp/PiOmxRemote .
	@echo "copy temp/PiOmxRemote to your pi and run"
