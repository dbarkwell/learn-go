package album

type Service struct {
	AlbumRepository Repository
}

func ProvideAlbumService(ar Repository) Service {
	return Service{AlbumRepository: ar}
}

func (as *Service) FindAll() ([]AlbumDTO, error) {
	albums, err := as.AlbumRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return ToAlbumDTOs(albums), nil
}

func (as *Service) FindByID(id int64) (AlbumDTO, error) {
	album, err := as.AlbumRepository.FindByID(id)
	if err != nil {
		return AlbumDTO{}, err
	}

	return ToAlbumDTO(album), nil
}

func (as *Service) Add(a NewAlbumDTO) (AlbumDTO, error) {
	newAlbum := ToNewAlbum(a)
	album, err := as.AlbumRepository.Add(newAlbum)
	if err != nil {
		return AlbumDTO{}, err
	}

	return ToAlbumDTO(album), nil
}

func (as *Service) Remove(id int64) (bool, error) {
	removed, err := as.AlbumRepository.Remove(id)
	if err != nil {
		return false, err
	}

	return removed, nil
}
