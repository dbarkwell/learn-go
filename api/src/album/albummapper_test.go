package album

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestToAlbumDTO(t *testing.T) {
	got := ToAlbumDTO(Album{ID: 1, Artist: "artist", Title: "title", Price: 10})
	want := AlbumDTO{ID: 1, Artist: "artist", Title: "title", Price: 10}

	if !cmp.Equal(got, want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}

func TestToAlbumDTOs(t *testing.T) {
	got := ToAlbumDTOs([]Album{
		{ID: 1, Artist: "artist1", Title: "title1", Price: 11},
		{ID: 2, Artist: "artist2", Title: "title2", Price: 12},
		{ID: 3, Artist: "artist3", Title: "title3", Price: 13},
	})

	want := []AlbumDTO{
		{ID: 1, Artist: "artist1", Title: "title1", Price: 11},
		{ID: 2, Artist: "artist2", Title: "title2", Price: 12},
		{ID: 3, Artist: "artist3", Title: "title3", Price: 13},
	}

	if !cmp.Equal(got, want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}
