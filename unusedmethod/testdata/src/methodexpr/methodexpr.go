package methodexpr

type Callback interface {
	Call(res any)
}

type Registry struct {
	callbacks []Callback
}

func (reg *Registry) Register(cb Callback) {
	reg.callbacks = append(reg.callbacks, cb)
}

func (reg *Registry) Notify(res any) {
	for _, cb := range reg.callbacks {
		// we have usage to the Callback.Call usage as expression
		Callback.Call(cb, res)
	}
}
