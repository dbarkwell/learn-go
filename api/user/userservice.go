package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
)

type Service struct {
	UserRepository Repository
	Cache          *memcache.Client
}

func ProvideUserService(ur Repository, mc *memcache.Client) Service {
	return Service{
		UserRepository: ur,
		Cache:          mc,
	}
}

func (s *Service) Add(username string, firstname string, lastname string, email string) (UserDTO, error) {
	user, err := s.UserRepository.Add(username, firstname, lastname, email)

	if err != nil {
		return UserDTO{}, err
	}

	return UserDTO{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil
}

func (s *Service) Get(username string) (UserDTO, error) {
	var user User
	u, err := s.Cache.Get(getUserKey(username))
	if err != nil {
		if errors.Is(err, memcache.ErrCacheMiss) {
			user, err = s.UserRepository.Get(username)
			b, _ := json.Marshal(user)
			s.Cache.Set(&memcache.Item{Key: getUserKey(username), Value: b})
		} else {
			return UserDTO{}, err
		}
	} else {
		json.Unmarshal(u.Value, &user)
	}

	return UserDTO{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil
}

func getUserKey(username string) string {
	return fmt.Sprintf("user_%s", username)
}
