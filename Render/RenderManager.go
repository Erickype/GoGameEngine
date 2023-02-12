package Render

import "github.com/Erickype/GoGameEngine/Common"

type Manager struct {
}

var instance *Manager

func (m Manager) New() Common.IManager {
	return &Manager{}
}

func (m Manager) GetInstance() Common.IManager {
	if instance == nil {
		instance = &Manager{}
	}
	return instance
}

func (m Manager) Start() Common.IManager {
	//TODO implement me
	panic("implement me")
}

func (m Manager) ShutDown() Common.IManager {
	//TODO implement me
	panic("implement me")
}
