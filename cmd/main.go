package main

import (
	"log"

	"github.com/sureshkumar-cse/SoundShift/internal/config"
	"github.com/sureshkumar-cse/SoundShift/internal/core"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Source and Destination Platforms
	sourcePlatform := "YouTube Music"
	destinationPlatform := "Amazon Music"
	playlistID := "YOUR_SOURCE_PLAYLIST_ID"

	// Fetch Songs from Source Playlist
	songs, err := core.FetchPlaylist(sourcePlatform, playlistID)
	if err != nil {
		log.Fatalf("Error fetching playlist: %v", err)
	}

	// Process each song
	for _, song := range songs {
		songID, err := core.SearchMusic(destinationPlatform, song)
		if err != nil {
			log.Printf("Song not found: %s - %s\n", song.Title, song.Artist)
			continue
		}

		err = core.AddToPlaylist(destinationPlatform, songID)
		if err != nil {
			log.Printf("Failed to add song %s to %s: %v\n", song.Title, destinationPlatform, err)
		}
	}
}
