package memorydb

import (
	"context"
	"errors"
	"sync"

	"github.com/pulse227/server-recruit-challenge-sample/model"
)

type albumRepository struct {
	sync.RWMutex
	albumMap map[model.AlbumID]*model.Album // キーが AlbumID、値が model.Album のマップ
}

func NewAlbumRepository() *albumRepository {
	var initMap = map[model.AlbumID]*model.Album{
		1: {ID: 1, Title: "Alice's 1st Album", SingerID: 1},
		2: {ID: 2, Title: "Alice's 2nd Album", SingerID: 1},
		3: {ID: 3, Title: "Bella's 1st Album", SingerID: 2},
	}

	return &albumRepository{
		albumMap: initMap,
	}
}

func (r *albumRepository) GetAll(ctx context.Context) ([]*model.Album, error) {
	r.RLock()
	defer r.RUnlock()

	albums := make([]*model.Album, 0, len(r.albumMap))
	for _, s := range r.albumMap {
		albums = append(albums, s)
	}
	return albums, nil
}

func (r *albumRepository) Get(ctx context.Context, id model.AlbumID) (*model.Album, error) {
	r.RLock()
	defer r.RUnlock()

	album, ok := r.albumMap[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return album, nil
}

func (r *albumRepository) Add(ctx context.Context, album *model.Album) error {
	r.Lock()
	r.albumMap[album.ID] = album
	r.Unlock()
	return nil
}

func (r *albumRepository) Delete(ctx context.Context, id model.AlbumID) error {
	r.Lock()
	delete(r.albumMap, id)
	r.Unlock()
	return nil
}

