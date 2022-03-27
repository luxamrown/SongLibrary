package delivery

import (
	"fmt"

	"enigmacamp.com/song/usecase"
)

func AddSongForm(usecase usecase.AddSongUseCase) {

	var songArtist string
	var songAlbum string
	var songTitle string

	fmt.Print("Artist: ")
	fmt.Scan(&songArtist)
	fmt.Print("Album: ")
	fmt.Scan(&songAlbum)
	fmt.Print("Title: ")
	fmt.Scan(&songTitle)
	usecase.Register(songArtist, songAlbum, songTitle)
	BackToMain()
}
