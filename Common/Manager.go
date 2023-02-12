package Common

type IManager interface {
	Init()
	GetInstance() IManager
	Start() IManager
	ShutDown() IManager
}
