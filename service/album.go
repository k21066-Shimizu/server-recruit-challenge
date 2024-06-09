package service

import (
	"context"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/repository"
)

type AlbumService interface {
	GetAlbumListService(ctx context.Context) ([]*model.Album, error)
	GetAlbumService(ctx context.Context, albumID model.AlbumID) (*model.AlbumWithSinger, error)
	PostAlbumService(ctx context.Context, album *model.Album) error
	DeleteAlbumService(ctx context.Context, albumID model.AlbumID) error
}

type albumService struct {
	albumRepository repository.AlbumRepository
	singerRepository repository.SingerRepository
}

var _ AlbumService = (*albumService)(nil)

func NewAlbumService(ar repository.AlbumRepository, sr repository.SingerRepository) *albumService {
	return &albumService{albumRepository: ar, singerRepository: sr}
}

func (s *albumService) GetAlbumListService(ctx context.Context) ([]*model.Album, error) {
	albums, err := s.albumRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return albums, nil
}

func (s *albumService) GetAlbumService(ctx context.Context, albumID model.AlbumID) (*model.AlbumWithSinger, error) {
	album, err := s.albumRepository.Get(ctx, albumID)
	if err != nil {
		return nil, err
	}

	singer, err := s.singerRepository.Get(ctx, album.SingerID)
	if err != nil {
		return nil, err
	}

	albumWithSinger := &model.AlbumWithSinger{
		ID:     album.ID,
		Title:  album.Title,
		Singer: *singer,
	}

	return albumWithSinger, nil
}

func (s *albumService) PostAlbumService(ctx context.Context, album *model.Album) error {
	if err := s.albumRepository.Add(ctx, album); err != nil {
		return err
	}
	return nil
}

func (s *albumService) DeleteAlbumService(ctx context.Context, albumID model.AlbumID) error {
	if err := s.albumRepository.Delete(ctx, albumID); err != nil {
		return err
	}
	return nil
}
