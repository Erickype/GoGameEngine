package Core

type _ interface {
	PushLayer(layer *iLayer)
	PushOverlay(overLay *iLayer)
	PopLayer(layer *iLayer)
	PopOverlay(overlay *iLayer)
}

type LayerStack struct {
	layers *[]*iLayer
}
