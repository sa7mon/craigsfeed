package data

import (
	"github.com/gorilla/feeds"
	"sync"
	"time"
)

type Error struct {
	Message string    `json:"error_message"`
	Time    time.Time `json:"error_time"`
}

type DataManager struct {
	CurrentFeed *feeds.Feed
	Lock        *sync.Mutex
	Error       *Error
}

func (dm *DataManager) SetError(err error) {
	dm.Lock.Lock()
	dm.Error = &Error{
		Message: err.Error(),
		Time:    time.Now(),
	}
	dm.Lock.Unlock()
}

func (dm *DataManager) ClearError() {
	dm.Lock.Lock()
	dm.Error = nil
	dm.Lock.Unlock()
}

func (dm *DataManager) GetError() *Error {
	dm.Lock.Lock()
	defer dm.Lock.Unlock()
	return dm.Error
}

var once sync.Once
var instance DataManager

func GetManager() *DataManager {
	once.Do(func() { // atomic, do only once
		instance = DataManager{Lock: &sync.Mutex{}}
	})

	return &instance
}
