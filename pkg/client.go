package pokego

type IPokeGo interface{}

type PokeGo struct {
	BaseUrl string
}

func NewClient() PokeGo {
	return PokeGo{
		BaseUrl: PokeApiBaseUrl,
	}
}
