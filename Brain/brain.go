package Brain

import "github.com/Erickype/GoGameEngine/Brain/BrainComponents"

type Brain struct {
	Name      string
	Forebrain BrainComponents.Forebrain
	Midbrain  BrainComponents.Midbrain
	Hindbrain BrainComponents.Hindbrain
}

func NewBrain(name string) Brain {
	return Brain{
		Name:      name,
		Forebrain: BrainComponents.Forebrain{},
		Midbrain:  BrainComponents.Midbrain{},
		Hindbrain: BrainComponents.Hindbrain{},
	}
}
