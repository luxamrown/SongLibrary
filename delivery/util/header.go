package util

import (
	"fmt"
	"strings"
)

func CreateHeader(title string) {
	fmt.Println(strings.Repeat("*", 75))
	fmt.Printf("%40s\n", title)
	fmt.Println(strings.Repeat("*", 75))
}

func CreateHeaderTable() {
	fmt.Printf(SongListTableHeaderFormat, NoLabel, SongArtistLabel, SongAlbumLabel, SongTitleLabel)
	fmt.Println(strings.Repeat("_", 75))
}
