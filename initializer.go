package initializer

import (
	"container/list"
	"log"
	"sync"
)

type Initializer = func() error
type Logger = func(string, ...interface{})

type initializerSet struct {
	name string
	init Initializer
}

var mux sync.Mutex
var setQueue = list.New()
// Function to log initializing process. Default `noOpLog`
var Log Logger = noOpLog

// Dummy log function, will not log anything
func noOpLog(format string, v ...interface{}) {}

// Register init function, the name must be unique
func Register(name string, init Initializer) {
	mux.Lock()
	defer mux.Unlock()
	setQueue.PushBack(&initializerSet{
		name: name,
		init: init,
	})
}

func Init() {
	mux.Lock()
	defer mux.Unlock()
	for setQueue.Len() > 0 {
		e := setQueue.Front()
		set := e.Value.(*initializerSet)
		Log("initializing %s", set.name)
		if err := set.init(); err != nil {
			format := "initializer (%s) encountered an error: %+v"
			Log(format, set.name, err)
			log.Panicf(format, set.name, err)
		}
		setQueue.Remove(e)
	}
}
