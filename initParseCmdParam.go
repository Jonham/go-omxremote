package main

import "flag"

var videosPath string
var bindAddr string
var omx string

func parseCmdParam() {
	flag.StringVar(&videosPath, "media", ".", "Path to look for videos in")
	flag.StringVar(&bindAddr, "bind", ":31415", "Address to bind on.")
	flag.StringVar(&omx, "omx", "-o hdmi", "Options to pass to omxplayer")
	flag.Parse()
}
