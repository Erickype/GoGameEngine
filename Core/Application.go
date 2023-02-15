package Core

import "log"

type IApplication interface {
	Run()
	Destroy()
}

type Application struct{}

func (Application) Run() {
	for {
		log.Println("Running app")
	}
}

func (Application) Destroy() {

}
