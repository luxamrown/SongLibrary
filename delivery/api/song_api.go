package api

import (
	"enigmacamp.com/song/delivery/appreq"
	"enigmacamp.com/song/manager"
	"enigmacamp.com/song/model"
	"github.com/gin-gonic/gin"
)

type SongApi struct {
	BaseApi
	usecase manager.UseCaseManager
}

func (s *SongApi) createSong() gin.HandlerFunc {
	return func(c *gin.Context) {
		var songReq appreq.SongRequest
		err := s.ParseRequestBody(c, &songReq)
		if err != nil {
			s.FailedRequest(c, "api-createSong", "02", "Can not parse body")
			return
		}
		// s.usecase.AddSongUseCase().Register(songReq.SongArtist, songReq.SongAlbum, songReq.SongTitle)
		SongRegistered, err := s.usecase.AddSongUseCase().Register(songReq.SongArtist, songReq.SongAlbum, songReq.SongTitle)
		if err != nil {
			s.Failed(c, "api-createSong", "01", "Can Not create product")
			return
		}
		s.Success(c, "Song", SongRegistered)
	}
}

func (s *SongApi) listSong() gin.HandlerFunc {
	return func(c *gin.Context) {
		var song []model.Song
		song, err := s.usecase.ShowSongListUseCase().ShowAll()
		if err != nil {
			c.JSON(404, gin.H{"error": "song not found"})
			return
		}
		c.JSON(200, song)

	}
}

func (s *SongApi) findSongByArtist() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var songReq appreq.SongRequest
		// err := s.ParseRequestBody(c, &songReq)
		// if err != nil {
		// 	s.FailedRequest(c, "api-createSong", "02", "Can not parse body")
		// 	return
		// }
		songArtist := c.Query("songArtist")
		selectedSong, err := s.usecase.FindSongByArtist().FindSongByArtist(songArtist)
		if err != nil {
			s.Failed(c, "api-findsong", "03", "Can Not find song")
		}
		c.JSON(200, selectedSong)
	}
}

func NewSongApi(songRoute *gin.RouterGroup, usecase manager.UseCaseManager) {
	api := SongApi{
		usecase: usecase,
	}
	songRoute.POST("", api.createSong())
	songRoute.GET("", api.listSong())
	songRoute.GET("/search", api.findSongByArtist())
}
