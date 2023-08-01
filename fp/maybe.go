package fp

type Maybe[T any] *T

func Some[T any](t T) Maybe[T] {
    return &t
}

func None[T any]() Maybe[T] {
    return nil
}

func IsNone[T any](v Maybe[T]) bool {
    return v == nil
}

func Just[T any](v Maybe[T]) T {
    return *v
}
