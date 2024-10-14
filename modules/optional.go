package modules

type Optional[T any] struct {
	OK    bool
	Value T
}
