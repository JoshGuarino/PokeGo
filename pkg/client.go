package pokego

func NewClient() PokeGo {
	return PokeGo{
		BaseUrl: PokeApiBaseUrl,
	}
}
