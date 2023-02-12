package Render

type Manager struct {
}

var instance *Manager

func (m *Manager) GetInstance() *Manager {
	if instance == nil {
		instance = &Manager{}
	}
	return instance
}

func (m *Manager) Start() {}

func (m *Manager) ShutDown() {}
