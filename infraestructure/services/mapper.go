package services

import (
	"github.com/alexis-aguirre/golang-bootcamp-2020/domain/model"
)

func lyricResponseToSong(lyricResponse *HappiLyricResponse) *model.Song {
	song := &model.Song{}
	song.ID = lyricResponse.Result.IDTrack
	song.Name = lyricResponse.Result.Track
	song.AlbumID = lyricResponse.Result.IDAlbum
	song.Album = lyricResponse.Result.Album
	song.Interpreter = lyricResponse.Result.Artist
	song.InterpreterID = lyricResponse.Result.IDArtist
	song.Lyric = lyricResponse.Result.Lyrics
	return song
}

func seachResponseToSongSlice(searchResponse *HappiSearchResponse) []*model.Song {
	var songs []*model.Song
	for i := 0; i < searchResponse.Length; i++ {
		currentResult := searchResponse.Result[i]
		song := &model.Song{}

		song.ID = currentResult.IDTrack
		song.Name = currentResult.Track
		song.AlbumID = currentResult.IDAlbum
		song.Album = currentResult.Album
		song.Interpreter = currentResult.Artist
		song.InterpreterID = currentResult.IDArtist
		if currentResult.HasLyrics {
			song.Lyric = "yes"
		} else {
			song.Lyric = "no"
		}

		songs = append(songs, song)
	}
	return songs
}
