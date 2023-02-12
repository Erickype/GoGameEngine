package Common

type IManager interface {
	New() IManager
	GetInstance() IManager
	Start() IManager
	ShutDown() IManager
}
