package api

import (
	"fmt"
)

// FetchYouTubePlaylist retrieves playlist from YouTube Music
func FetchYouTubePlaylist(playlistID string) ([]map[string]string, error) {
	fmt.Println("Fetching YouTube Music playlist:", playlistID)
	// Mock response
	return []map[string]string{
		{"title": "Song A", "artist": "Artist A", "album": "Album A"},
		{"title": "Song B", "artist": "Artist B", "album": "Album B"},
	}, nil
}
