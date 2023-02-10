package Brain

import "github.com/Erickype/GoGameEngine/Brain/BrainComponents"

type Brain struct {
	Name      string
	Forebrain BrainComponents.Forebrain
	Midbrain  BrainComponents.Midbrain
	Hindbrain BrainComponents.Hindbrain
}
