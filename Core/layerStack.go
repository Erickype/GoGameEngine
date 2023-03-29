package Core

type _ interface {
	PushLayer(layer *ILayer)
	PushOverlay(overLay *ILayer)
	PopLayer(layer *ILayer)
	PopOverlay(overlay *ILayer)
	Construct()
	Begin() *ILayer
	End() *ILayer
}

type LayerStack struct {
	layers      *[]*ILayer
	layerInsert *int
}

func (l *LayerStack) PushLayer(layer *ILayer) {
	*l.layers = append(*l.layers, nil)
	copy((*l.layers)[*l.layerInsert+1:], (*l.layers)[*l.layerInsert:])
	(*l.layers)[*l.layerInsert] = layer
	*l.layerInsert++
}

func (l *LayerStack) PushOverlay(overLay *ILayer) {
	*l.layers = append(*l.layers, overLay)
}

func (l *LayerStack) PopLayer(layer *ILayer) {
	for i, ly := range *l.layers {
		if ly == layer {
			*l.layers = append((*l.layers)[:i], (*l.layers)[i+1:]...)
			*l.layerInsert--
			break
		}
	}
}

func (l *LayerStack) PopOverlay(overlay *ILayer) {
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
	layers := make([]*ILayer, 0)
	l.layers = &layers
}

func (l *LayerStack) Begin() *ILayer {
	return (*l.layers)[0]
}

func (l *LayerStack) End() *ILayer {
	return (*l.layers)[len(*l.layers)-1]
}
