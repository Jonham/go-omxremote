all:omx

assets:
	esc -o clientAssets.go -prefix="dist/app" dist

omx:assets
	go build -o temp/OMXRemote .

prepare:assets
	echo "now you can run `go run .`"
