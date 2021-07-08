package watcher

import (
	"log"

	fsn "github.com/fsnotify/fsnotify"
)

const (
	welcome = "Filesystem Notify Workshop"
)

// New watcher file system
func New() *Watcher {
	return &Watcher{}
}

// Watcher data collection
type Watcher struct {
}

// Service is interface for watcher method
type Service interface {
	Run()
}

// Run watcher package
func (w *Watcher) Run() {
	log.Println(welcome)

	watcher, err := fsn.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				log.Println("Event:", event)
				log.Println("Name:", event.Name)
				log.Println("Operation:", event.Op)
				if event.Op&fsn.Write == fsn.Write {
					log.Println("Modified file:", event.Name)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}

				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("./watch_me")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
