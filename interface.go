package main

// Interface: PlaySong
type PlaySong interface {
	SongPlay(songs string) string
	SongStop() string
}

// Struct: SongList
type SongList struct{}

// Method: SongPlay (implements PlaySong interface)
func (s SongList) SongPlay(song string) string {
	return "Now Playing: " + song
}

// Method SongStop (implements PlaySong interface)
func (s SongList) SongStop() string {
	return "Song Stopped"
}

/*

func main() {
	// Interface type variable
	var player PlaySong

	// SongList struct assign করলাম interface variable এ
	player = SongList{}

	// Function call
	fmt.Println(player.SongPlay("Tum Hi Ho"))
	fmt.Println(player.SongStop())
}

*/
