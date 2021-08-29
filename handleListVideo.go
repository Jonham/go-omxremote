package main

import (
	"encoding/base64"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/karrick/godirwalk"
	"net/http"
	"path/filepath"
)

// Video struct contains has two fields:
// filename and base32 hash of the filepath
type Video struct {
	File string `json:"file"`
	Hash string `json:"hash"`
}

// List function - outputs json with all video files in the videoPath
func handleListVideo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var files []*Video
	var root = videosPath

	addCorsHeader(w)

	_ = godirwalk.Walk(root, &godirwalk.Options{
		Unsorted: true, // set true for faster yet non-deterministic enumeration (see godoc)
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			if de.IsDir() == false {
				if filepath.Ext(osPathname) == ".mkv" || filepath.Ext(osPathname) == ".mp4" || filepath.Ext(osPathname) == ".avi" || filepath.Ext(osPathname) == ".mov" {
					files = append(files, &Video{File: filepath.Base(osPathname), Hash: base64.URLEncoding.EncodeToString([]byte(osPathname))})
				}
			}
			return nil
		},
		ErrorCallback: func(osPathname string, err error) godirwalk.ErrorAction {
			return godirwalk.SkipNode
		},
	})

	encoder := json.NewEncoder(w)
	encoder.Encode(files)
}
