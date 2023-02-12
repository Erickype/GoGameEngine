package Common

type IManager interface {
	GetInstance()
	Start()
	ShutDown()
}
