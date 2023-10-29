package album

func ToAlbum(dto AlbumDTO) Album {
	return Album{ID: dto.ID, Title: dto.Title, Artist: dto.Artist, Price: dto.Price}
}

func ToNewAlbum(dto NewAlbumDTO) Album {
	return Album{Title: dto.Title, Artist: dto.Artist, Price: dto.Price}
}

func ToAlbumDTO(a Album) AlbumDTO {
	return AlbumDTO{ID: a.ID, Title: a.Title, Artist: a.Artist, Price: a.Price}
}

func ToAlbumDTOs(a []Album) []AlbumDTO {
	albumDTOs := make([]AlbumDTO, len(a))

	for i, itm := range a {
		albumDTOs[i] = ToAlbumDTO(itm)
	}

	return albumDTOs
}
