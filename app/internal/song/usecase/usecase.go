package usecase

import (
	client "effective_project/app/internal/client"
	songRep "effective_project/app/internal/song/repository"
	"effective_project/app/models"
	"strings"

	"github.com/sirupsen/logrus"
)

type useCase struct {
	songRepository songRep.RepositoryI
	client         client.Client
}

type UseCaseI interface {
	GetSongLib(songParam *models.Song, limit, offset int) ([]*models.Song, error)
	GetSongText(songId string, limit, offset int) (*[]string, error)
	DeleteSong(songId string) error
	EditSong(song *models.Song) error
	CreateSong(song *models.Song) error
}

func New(songRepository songRep.RepositoryI, client client.Client) UseCaseI {
	return &useCase{
		songRepository: songRepository,
		client:         client,
	}
}

func (uc *useCase) GetSongLib(songParam *models.Song, limit, offset int) ([]*models.Song, error) {
	logrus.WithFields(logrus.Fields{
		"songGroup":       songParam.Group,
		"songReleaseDate": songParam.ReleaseDate,
		"songTitle":       songParam.Title,
		"limit":           limit,
		"offset":          offset,
	}).Debug("Deserialized data from client")
	songs, err := uc.songRepository.SelectSongs(songParam, limit, offset)
	if err != nil {
		return nil, models.ErrBadData
	}
	logrus.Info("Data succesfully received from DB")
	return songs, nil
}

func (uc *useCase) GetSongText(songId string, limit, offset int) (*[]string, error) {
	logrus.WithFields(logrus.Fields{
		"songId": songId,
		"limit":  limit,
		"offset": offset,
	}).Debug("Deserialized data from client")
	song, err := uc.songRepository.SelectSongById(songId)
	if err != nil {
		return nil, models.ErrSongNotFound
	}
	song.Text = strings.ReplaceAll(song.Text, `\n`, "\n")
	verses := strings.Split(song.Text, "\n\n")
	if offset < len(verses) {
		verses = verses[offset:]
	}
	if limit > 0 && limit < len(verses) {
		verses = verses[:limit]
	}
	logrus.WithFields(logrus.Fields{
		"songText": song.Text,
	}).Debug("Data from DB")
	return &verses, nil
}

func (uc *useCase) DeleteSong(songId string) error {
	logrus.WithFields(logrus.Fields{
		"songId": songId,
	}).Debug("Deserialized data from client")
	_, err := uc.songRepository.SelectSongById(songId)
	if err != nil {
		return models.ErrSongNotFound
	}
	err = uc.songRepository.DeleteSongById(songId)
	if err == nil {
		logrus.Info("Song succesfully deleted from DB")
	}
	return err
}

func (uc *useCase) EditSong(song *models.Song) error {
	logrus.WithFields(logrus.Fields{
		"songId":          song.Id,
		"songTitle":       song.Title,
		"songGroup":       song.Group,
		"songLink":        song.Link,
		"songText":        song.Text,
		"songReleaseDate": song.ReleaseDate,
	}).Debug("Deserialized data from client")
	_, err := uc.songRepository.SelectSongById(song.Id)
	if err != nil {
		return models.ErrSongNotFound
	}
	err = uc.songRepository.EditSongById(song)
	if err == nil {
		logrus.Info("Song succesfully edited")
	}
	return err
}

func (uc *useCase) CreateSong(song *models.Song) error {
	logrus.WithFields(logrus.Fields{
		"songTitle": song.Title,
		"songGroup": song.Group,
	}).Debug("Deserialized data from client")
	songInfo, err := uc.client.FetchSongInfo(song.Group, song.Title)
	if err != nil {
		return models.ErrInternalServer
	}

	logrus.WithFields(logrus.Fields{
		"songLink":        song.Link,
		"songText":        song.Text,
		"songReleaseDate": song.ReleaseDate,
	}).Debug("Deserialized data from serverInfoSongs")
	song.ReleaseDate = songInfo.ReleaseDate
	song.Link = songInfo.Link
	song.Text = songInfo.Text
	err = uc.songRepository.CreateSong(song)
	if err == nil {
		logrus.Info("Song succesfully created")
	}
	return err
}
