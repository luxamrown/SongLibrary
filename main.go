package main

import (
	"fmt"

	"enigmacamp.com/song/config"
	"enigmacamp.com/song/delivery"
)

func main() {
	appConfig := config.NewConfig()
	delivery.MainMenu()
	for {
		var choiceMain string
		fmt.Scan(&choiceMain)
		switch choiceMain {
		case "1":
			delivery.AddSongForm(appConfig.UseCaseManager.AddSongUseCase())
		case "2":
			delivery.ShowSongList(appConfig.UseCaseManager.ShowSongListUseCase())
		case "3":
			delivery.FindSong(appConfig.UseCaseManager.FindSongByArtist())
		case "4":
			delivery.ExitApp()
		}
	}
}
