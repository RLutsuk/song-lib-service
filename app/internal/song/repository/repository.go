package repository

import (
	"effective_project/app/models"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type dataBase struct {
	db *gorm.DB
}

type RepositoryI interface {
	SelectSongs(songParam *models.Song, limit, offset int) ([]*models.Song, error)
	SelectSongById(songId string) (*models.Song, error)
	DeleteSongById(songId string) error
	EditSongById(song *models.Song) error
	CreateSong(song *models.Song) error
}

func New(db *gorm.DB) RepositoryI {
	return &dataBase{
		db: db,
	}
}

func (dbSong *dataBase) SelectSongs(songParam *models.Song, limit, offset int) ([]*models.Song, error) {
	songs := make([]*models.Song, 0, limit)
	tx := dbSong.db.Table("songs")

	if songParam.Title != "" {
		tx = tx.Where("title ILIKE ?", "%"+songParam.Title+"%")
	}

	if songParam.Group != "" {
		tx = tx.Where("group_song ILIKE ?", "%"+songParam.Group+"%")
	}

	if !songParam.ReleaseDate.IsZero() {
		tx = tx.Where("release_date = ?", songParam.ReleaseDate)
	}
	tx = tx.Limit(limit).Offset(offset).Order("title ASC").Select("title", "group_song", "release_date", "link_song")

	err := tx.Find(&songs).Error
	if err != nil {
		return nil, err
	}

	return songs, nil
}

func (dbSong *dataBase) SelectSongById(songId string) (*models.Song, error) {
	song := models.Song{}
	tx := dbSong.db.Table("songs").Where("id = ?", songId).Take(&song)
	if tx.Error != nil {
		return &song, errors.Wrap(tx.Error, "database error (table songs)")
	}
	return &song, nil
}

func (dbSong *dataBase) DeleteSongById(songId string) error {
	song := models.Song{}
	tx := dbSong.db.Table("songs").Where("id = ?", songId).Delete(song)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table songs)")
	}
	return nil
}

func (dbSong *dataBase) EditSongById(song *models.Song) error {
	tx := dbSong.db.Table("songs").Where("id = ?", song.Id).Updates(*song)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table songs)")
	}
	return nil
}

func (dbSong *dataBase) CreateSong(song *models.Song) error {
	tx := dbSong.db.Table("songs").Create(song)
	if tx.Error != nil {
		return errors.Wrap(tx.Error, "database error (table tenders)")
	}
	return nil
}
