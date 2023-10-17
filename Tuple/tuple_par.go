package Tuple

// Transform both parts of a tuple.
func MapBoth_par[T, U, E, D any](mapfnF func(valueF T) E, mapfnS func(valueS U) D, t Tuple[T, U]) Tuple[E, D] {
	chE := make(chan E)
	chD := make(chan D)
	go func() {
		chE <- mapfnF(t.Fst)
	}()
	go func() {
		chD <- mapfnS(t.Snd)
	}()
	return Tuple[E, D]{<-chE, <-chD}
}
