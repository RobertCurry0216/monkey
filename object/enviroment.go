package object

//NewEnviroment creates a new enviroment
func NewEnviroment() *Enviroment {
	s := make(map[string]Object)
	return &Enviroment{store: s}
}

//NewEnclosedEnviroment creates a new enviroment with a pointer to the given Enviroment
func NewEnclosedEnviroment(outer *Enviroment) *Enviroment {
	env := NewEnviroment()
	env.outer = outer
	return env
}

// Enviroment is a container for the vairables
type Enviroment struct {
	store map[string]Object
	outer *Enviroment
}

// Get returns a variable object from its name
func (e *Enviroment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

//Set adds an object to the enviroment store
func (e *Enviroment) Set(name string, obj Object) Object {
	e.store[name] = obj
	return obj
}
