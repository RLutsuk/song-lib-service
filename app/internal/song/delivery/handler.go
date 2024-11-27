package delivery

import (
	songUC "effective_project/app/internal/song/usecase"
	"effective_project/app/models"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type delivery struct {
	songUC songUC.UseCaseI
}

// @Summary      Получение списка песен
// @Description  Получение списка из 20 песен упорядоченного по названию с пагинацией.
// @Description  Есть возможность филтрации по дате релиза, названию группы, названию песни, а также возможность изменения количества песен на странице.
// @Tags         API Manage songs
// @Accept       json
// @Produce      json
// @Param        group        query    string  false "Group"
// @Param        title        query    string  false "Title"
// @Param        releaseDate  query    string  false "Release Date"
// @Param        limit        query    int     false "Limit"  default(20)
// @Param        offset       query    int     false "Offset" default(0)
// @Success      200          {array}  models.Song "List of songs"
// @Failure      400          {object} string     "Bad request" SchemaExample({"code":400, "message":"Bad request"})
// @Failure      500          {object} string     "Internal server error" SchemaExample({"code":500, "message":"Internal server error"})
// @Router       /songs [get]
func (delivery *delivery) getSongLibrary(c echo.Context) error {
	logrus.Info("Start processing the request: getting library of songs")
	var songParam models.Song
	songParam.Group = c.QueryParam("group")
	songParam.Title = c.QueryParam("title")

	layout := "02.01.2006"
	if c.QueryParam("releaseDate") != "" {
		parsedDate, err := time.Parse(layout, c.QueryParam("releaseDate"))
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadData.Error())
		}
		songParam.ReleaseDate = parsedDate
	}

	limitStr := c.QueryParam("limit")
	offsetStr := c.QueryParam("offset")
	limit := 20
	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadData.Error())
		}
		if limit < 10 || limit > 100 {
			limit = 20
		}
	}

	offset := 0
	if offsetStr != "" {
		var err error
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadData.Error())
		}
		if offset < 0 {
			offset = 0
		}
	}

	songs, err := delivery.songUC.GetSongLib(&songParam, limit, offset)

	if err != nil {
		switch {
		case errors.Is(err, models.ErrBadData):
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadData.Error())
		default:
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServer.Error())
		}
	}
	logrus.Info("End of process the request: getting library of songs")
	return c.JSON(http.StatusOK, songs)
}

// @Summary      Получение текста песни
// @Description  Получение текста песни по ID с пагинацией по куплетам.
// @Tags         API Manage songs
// @Accept       json
// @Produce      json
// @Param        songId  path      string  true  "Song ID"
// @Param        limit   query     int     false "Limit"  default(10)
// @Param        offset  query     int     false "Offset" default(0)
// @Success      200     {array}   string  "Song lyrics text"         SchemaExample({"code":200, "message":"Textttt of song"})
// @Failure      404     {object}  string  "Song not found"           SchemaExample({"code":404, "message":"Song not found"})
// @Failure      500     {object}  string  "Internal server error"    SchemaExample({"code":500, "message":"Internal server error."})
// @Router       /songs/{songId}/textsong [get]
func (delivery *delivery) getSongText(c echo.Context) error {
	logrus.Info("Start processing the request: getting text of song")
	idSong := c.Param("songId")
	limitStr := c.QueryParam("limit")
	offsetStr := c.QueryParam("offset")
	limit := 10
	if limitStr != "" {
		var err error
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadData.Error())
		}
		if limit < 0 {
			limit = 10
		}
	}

	offset := 0
	if offsetStr != "" {
		var err error
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadData.Error())
		}
		if offset < 0 {
			offset = 0
		}
	}

	text, err := delivery.songUC.GetSongText(idSong, limit, offset)

	if err != nil {
		switch {
		case errors.Is(err, models.ErrSongNotFound):
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusNotFound, models.ErrSongNotFound.Error())
		default:
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServer.Error())
		}
	}
	logrus.Info("End of process the request: getting text of song")
	return c.JSON(http.StatusOK, text)
}

// @Summary      Удаление песни
// @Description  Удаление песни из библиотеки по ID.
// @Tags         API Manage songs
// @Accept       json
// @Produce      json
// @Param        songId  path      string  true  "Song ID"
// @Success      200     {string}  string  "Success message"          SchemaExample({"code":200, "message":"The song was deleted from the library."})
// @Failure      404     {object}  string  "Song not found"           SchemaExample({"code":404, "message":"Song not found"})
// @Failure      500     {object}  string  "Internal server error"    SchemaExample({"code":500, "message":"Internal server error."})
func (delivery *delivery) deleteSong(c echo.Context) error {
	logrus.Info("Start processing the request: remove songs")
	idSong := c.Param("songId")

	err := delivery.songUC.DeleteSong(idSong)

	if err != nil {
		switch {
		case errors.Is(err, models.ErrSongNotFound):
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusNotFound, models.ErrSongNotFound.Error())
		default:
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServer.Error())
		}
	}
	logrus.Info("End of process the request: remove songs")
	return c.JSON(http.StatusOK, "The song was deleted from the library.")
}

// @Summary      Редактирование песни
// @Description  Редактирование параметров песни по Id. Все поля песни являются редактируемыми.
// @Tags         API Manage songs
// @Accept       json
// @Produce      json
// @Param        songId  path      string       true  "Song ID"
// @Param        song    body      models.Song  false "Song details to edit"
// @Success      200     {string}  string       "Success message"               SchemaExample({"code":200, "message":"The song has been successfully edited."})
// @Failure      400     {object}  string       "Bad request"                   SchemaExample({"code":400, "message":"Bad request"})
// @Failure      404     {object}  string       "Song not found"                SchemaExample({"code":404, "message":"Song not found"})
// @Failure      500     {object}  string       "Internal server error"         SchemaExample({"code":500, "message":"Internal server error."})
// @Router       /songs/{songId}/editsong [put]
func (delivery *delivery) editSong(c echo.Context) error {
	logrus.Info("Start processing the request: editing songs")
	var song models.Song
	if err := c.Bind(&song); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}
	song.Id = c.Param("songId")

	err := delivery.songUC.EditSong(&song)

	if err != nil {
		switch {
		case errors.Is(err, models.ErrSongNotFound):
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusNotFound, models.ErrSongNotFound.Error())
		default:
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	logrus.Info("End of process the request: editing songs")
	return c.JSON(http.StatusOK, "The song has been successfully edited.")
}

// @Summary      Добавление песни
// @Description  Добавление песни в библиотеку. При добавлении записываются только название песни и группы. Далее происходит запрос в сторонний сервис и после песня добавляется в библиотеку.
// @Tags         API Manage songs
// @Accept       json
// @Produce      json
// @Param        song    body      models.Song  true  "Song details to create"
// @Success      200     {string}  string       "Success message"                  	SchemaExample({"code":200, "message":"The song has been added to the library."})
// @Failure      400     {object}  string       "Bad request or missing required fields" SchemaExample({"code":400, "message":"Bad request"})
// @Failure      500     {object}  string       "Internal server error"            SchemaExample({"code":500, "message":"Internal server error."})
// @Router       /songs/newsong [post]
func (delivery *delivery) createSong(c echo.Context) error {
	logrus.Info("Start processing the request: creating songs")
	var song models.Song
	if err := c.Bind(&song); err != nil || song.Group == "" || song.Title == "" {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, models.ErrBadData)
	}

	err := delivery.songUC.CreateSong(&song)

	if err != nil {
		switch {
		case errors.Is(err, models.ErrSongNotFound):
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusNotFound, models.ErrSongNotFound.Error())
		default:
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, models.ErrInternalServer.Error())
		}
	}
	logrus.Info("End of process the request: creating songs")
	return c.JSON(http.StatusOK, "The song has been added to the library.")
}

func NewDelivery(e *echo.Echo, songUC songUC.UseCaseI) {
	handler := &delivery{
		songUC: songUC,
	}
	e.GET("/songs", handler.getSongLibrary)
	e.GET("/songs/:songId/textsong", handler.getSongText)
	e.DELETE("/songs/:songId/deletesong", handler.deleteSong)
	e.PATCH("/songs/:songId/editsong", handler.editSong)
	e.POST("/songs/newsong", handler.createSong)
	logrus.Info("Creation of processors")
}
