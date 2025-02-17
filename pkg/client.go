package pokego

type IPokeGo interface {
}

type PokeGo struct {
}

func NewClient() PokeGo {
	return PokeGo{}
}
