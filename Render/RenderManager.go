package Render

import (
	"github.com/Erickype/GoGameEngine/Common"
	"github.com/Erickype/GoGameEngine/Texture"
)

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
	Texture.Manager{}.Init()
	return instance
}

func (Manager) ShutDown() Common.IManager {
	//TODO implement me
	panic("implement me")
}
