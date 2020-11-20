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
