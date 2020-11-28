package object

import "fmt"

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

type Environment struct {
	store map[string]Object
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	//fmt.Printf("name: %s type %T\n", name, val)
	e.store[name] = val
	return val
}

//For debugging help
func (e *Environment) Print() {
	fmt.Println(e.store)
}
