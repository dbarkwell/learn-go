package user

type API struct {
	UserService Service
}

func ProvideUserAPI(us Service) API {
	return API{UserService: us}
}
