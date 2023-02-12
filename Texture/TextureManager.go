package Texture

import "github.com/Erickype/GoGameEngine/Common"

type Manager struct{}

var instance *Manager

func (m Manager) Init() {
	m.GetInstance()
	m.Start()
}

func (m Manager) GetInstance() Common.IManager {
	if instance == nil {
		instance = &Manager{}
	}
	return instance
}

func (m Manager) Start() Common.IManager {
	return instance
}

func (m Manager) ShutDown() Common.IManager {
	//TODO implement me
	panic("implement me")
}
