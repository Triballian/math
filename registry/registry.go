package registry

type registry struct {
	state *State
	
}

var instance *registry

func Get() *registery {
	if instance == nil {
		instance = &registry
	}
	return instance
}

func (this *registry) GetState() *State {
	
}