package authentication

type Service struct {
	AuthenticationRepository Repository
}

func ProvideAuthenticationService(ar Repository) Service {
	return Service{AuthenticationRepository: ar}
}
