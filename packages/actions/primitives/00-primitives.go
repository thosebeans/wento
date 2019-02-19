package primitives

func init() {
    primitives = map[string](func([]string)(Primitive, error)){}
}

var primitives map[string]func([]string)(Primitive, error)

func initPrimitives() {
    if primitives == nil {
        primitives = map[string](func([]string)(Primitive,error)){}
    }
}
