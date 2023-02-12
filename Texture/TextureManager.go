package Texture

import (
	"github.com/Erickype/GoGameEngine/Common"
)

type Manager struct{}

var instance *Manager

func (m Manager) New() Common.IManager {
	//TODO implement me
	panic("implement me")
}

func (Manager) GetInstance() Common.IManager {
	if instance == nil {
		instance = &Manager{}
	}
	return instance
}

func (Manager) Start() Common.IManager {
	//TODO implement me
	panic("implement me")
}

func (Manager) ShutDown() Common.IManager {
	//TODO implement me
	panic("implement me")
}
