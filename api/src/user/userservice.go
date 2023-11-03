package user

type Service struct {
	UserRepository Repository
}

func ProvideUserService(ur Repository) Service {
	return Service{UserRepository: ur}
}
