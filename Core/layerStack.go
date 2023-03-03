package Core

type _ interface {
	PushLayer(layer *Layer)
	PushOverlay(overLay *Layer)
	PopLayer(layer *Layer)
	PopOverlay(overlay *Layer)
	Construct()
}

type LayerStack struct {
	layers      *[]*Layer
	layerInsert *int
}

func (l *LayerStack) PushLayer(layer *Layer) {
	*l.layers = append((*l.layers)[:*l.layerInsert], layer)
	*l.layerInsert++
}

func (l *LayerStack) PushOverlay(overLay *Layer) {
	*l.layers = append(*l.layers, overLay)
}

func (l *LayerStack) PopLayer(layer *Layer) {
	for i, ly := range *l.layers {
		if ly == layer {
			*l.layers = append((*l.layers)[:i], (*l.layers)[i+1:]...)
			*l.layerInsert--
			break
		}
	}
}

func (l *LayerStack) PopOverlay(overlay *Layer) {
	for i := len(*l.layers) - 1; i >= *l.layerInsert; i-- {
		ly := (*l.layers)[i]
		if ly == overlay {
			*l.layers = append((*l.layers)[:i], (*l.layers)[i+1:]...)
			break
		}
	}
}

func (l *LayerStack) Construct() {
	i := 0
	l.layerInsert = &i
	layers := make([]*Layer, 1)
	l.layers = &layers
}
