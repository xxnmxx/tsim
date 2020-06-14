package tsim

type Enviroment struct {
	store map[string]Object
	outer *Enviroment
}

func NewEnviroment() *Enviroment {
	s := make(map[string]Object)
	return &Enviroment{store: s, outer: nil}
}

func (e *Enviroment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Enviroment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
