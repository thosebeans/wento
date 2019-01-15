package primitives

type Primitive interface {
    Test() error
    Emerge() error
}
