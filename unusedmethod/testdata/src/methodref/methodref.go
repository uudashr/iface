package methodref

type Callback interface {
	Call(res any)
}

type Registry struct {
	callbacks []func(res any)
}

func (reg *Registry) Register(cb Callback) {
	// we have usage to the Callback.Call as reference
	reg.callbacks = append(reg.callbacks, cb.Call)
}

func (reg *Registry) Notify(res any) {
	for _, cb := range reg.callbacks {
		cb(res)
	}
}
