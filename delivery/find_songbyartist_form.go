package delivery

import (
	"fmt"

	"enigmacamp.com/song/delivery/util"
	"enigmacamp.com/song/usecase"
)

func FindSong(usecase usecase.FindSongByArtist) {
	var songArtist string
	fmt.Print("Artist: ")
	fmt.Scan(&songArtist)

	song := usecase.FindSongByArtist(songArtist)
	if song != nil {
		util.CreateHeaderTable()
		for idx, s := range song {
			fmt.Printf(util.SongListTableFormat, idx+1, s.SongArtist, s.SongAlbum, s.SongTitle)
		}
	} else {
		fmt.Println("Produk tidak ditemukan")
	}
	BackToMain()
}
