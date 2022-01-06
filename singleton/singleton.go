package singleton

import (
	"fmt"
	"sync"
)

// Manager is the struct that we promote as a Singleton
type Manager struct{}

var lock = &sync.Mutex{}

var singleInstance *Manager

func GetManagerInstance() *Manager {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating Singleton Manager Now")
			singleInstance = &Manager{}
		}
	}
	return singleInstance
}

func (m *Manager) DoWork() error {
	fmt.Println("Singleton is doing some work")
	return nil
}
