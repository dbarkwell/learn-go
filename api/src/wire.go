//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"learn-go.barkwell.com/album"
	auth "learn-go.barkwell.com/authentication"
	"learn-go.barkwell.com/user"
)

func initAlbumAPI(db *sqlx.DB) album.API {
	wire.Build(album.ProvideAlbumRepository, album.ProvideAlbumService, album.ProvideAlbumAPI)

	return album.API{}
}

func initUserAPI(db *sqlx.DB) user.API {
	wire.Build(user.ProvideUserRepository, user.ProvideUserService, user.ProvideUserAPI)

	return user.API{}
}

func initAuthenticationAPI(db *sqlx.DB) auth.API {
	wire.Build(auth.ProvideAuthenticationRepository, auth.ProvideAuthenticationService, auth.ProvideAuthenticationAPI)

	return auth.API{}
}
