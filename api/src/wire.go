//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"learn-go.barkwell.com/album"
)

func initAlbumAPI(db *sqlx.DB) album.API {
	wire.Build(album.ProvideAlbumRepository, album.ProvideAlbumService, album.ProvideAlbumAPI)

	return album.API{}
}
