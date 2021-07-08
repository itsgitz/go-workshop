package main

import (
	"github.com/itsgitz/go-workshop/advanced/fsnotify/simple/internal/watcher"
)

func main() {
	watcher := watcher.New()
	watcher.Run()
}
