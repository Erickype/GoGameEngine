package Core

import "log"

type IApplication interface {
	run()
	destroy()
}

type Application struct{}

func (Application) run() {
	for {
		log.Println("Running app")
	}
}

func (Application) destroy() {
}

// CreateApplication This is the entry point to create an application
func CreateApplication() {
	log.Println("Starting engine")
	application := Application{}
	application.run()
	application.destroy()
}
