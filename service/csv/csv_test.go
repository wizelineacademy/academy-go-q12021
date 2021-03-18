package csvservice

import (
	"net/http"
	"os"
	"pokeapi/model"
	"reflect"
	"testing"
)

var pokemonsTest = []model.Pokemon{
	{ID: 1, Name: "greninja", URL: "https://pokeapi.co/api/v2/pokemon/658/"},
	{ID: 2, Name: "ursaring", URL: "https://pokeapi.co/api/v2/pokemon/217/"},
	{ID: 3, Name: "arcanine", URL: "https://pokeapi.co/api/v2/pokemon/59/"},
	{ID: 4, Name: "gengar", URL: "https://pokeapi.co/api/v2/pokemon/94/"},
	{ID: 5, Name: "porygon", URL: "https://pokeapi.co/api/v2/pokemon/137/"},
	{ID: 6, Name: "flareon", URL: "https://pokeapi.co/api/v2/pokemon/136/"},
	{ID: 7, Name: "omanyte", URL: "https://pokeapi.co/api/v2/pokemon/138/"},
	{ID: 8, Name: "frillish", URL: "https://pokeapi.co/api/v2/pokemon/592/"},
	{ID: 9, Name: "cacturne", URL: "https://pokeapi.co/api/v2/pokemon/332/"},
	{ID: 10, Name: "scizor", URL: "https://pokeapi.co/api/v2/pokemon/212/"},
}

var pokemonsTestLines = [][]string{
	{"1", "greninja", "https://pokeapi.co/api/v2/pokemon/658/"},
	{"2", "ursaring", "https://pokeapi.co/api/v2/pokemon/217/"},
	{"3", "arcanine", "https://pokeapi.co/api/v2/pokemon/59/"},
	{"4", "gengar", "https://pokeapi.co/api/v2/pokemon/94/"},
	{"5", "porygon", "https://pokeapi.co/api/v2/pokemon/137/"},
	{"6", "flareon", "https://pokeapi.co/api/v2/pokemon/136/"},
	{"7", "omanyte", "https://pokeapi.co/api/v2/pokemon/138/"},
	{"8", "frillish", "https://pokeapi.co/api/v2/pokemon/592/"},
	{"9", "cacturne", "https://pokeapi.co/api/v2/pokemon/332/"},
	{"10", "scizor", "https://pokeapi.co/api/v2/pokemon/212/"},
}

var pokemonsFromHttp = &[]model.SinglePokeExternal{
	{Name: "delcatty", URL: "https://pokeapi.co/api/v2/pokemon/301/"},
	{Name: "sableye", URL: "https://pokeapi.co/api/v2/pokemon/302/"},
	{Name: "mawile", URL: "https://pokeapi.co/api/v2/pokemon/303/"},
	{Name: "aron", URL: "https://pokeapi.co/api/v2/pokemon/304/"},
	{Name: "lairon", URL: "https://pokeapi.co/api/v2/pokemon/305/"},
}

func TestOpen(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "success openning",
			path:    "./file/pokemon.csv",
			wantErr: false,
		},
		{
			name:    "unsuccess openning",
			path:    "./file/pokemon2.csv",
			wantErr: true,
		},
	}
	const unsuccess string = "unsuccess openning"
	const success string = "success openning"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Open(tt.path)
			if (err != nil) != tt.wantErr && tt.name == success {
				t.Errorf("Open() error = %v, path %v", err, tt.path)
				return
			}
			if (err != nil) != tt.wantErr && tt.name == unsuccess {
				t.Errorf("Open() error = %v, path %v", err, tt.path)
				return
			}
		})
	}
}

func TestOpenAndWrite(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "success openning",
			path:    "./file/pokemon.csv",
			wantErr: false,
		},
		{
			name:    "unsuccess openning",
			path:    "./file/pokemon2.csv",
			wantErr: true,
		},
	}
	const unsuccess string = "unsuccess openning"
	const success string = "success openning"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := OpenAndWrite(tt.path)
			if (err != nil) != tt.wantErr && tt.name == success {
				t.Errorf("Open() error = %v, path %v", err, tt.path)
				return
			}
			if (err != nil) == tt.wantErr && tt.name == unsuccess {
				t.Errorf("Open() error = %v, path %v", err, tt.path)
				return
			}
		})
	}
}

func TestRead(t *testing.T) {
	f, _ := Open("./file/pokemon.csv")

	tests := []struct {
		name    string
		file    *os.File
		want    []model.Pokemon
		wantErr *model.Error
	}{
		{
			name:    "success reading pokemons",
			file:    f,
			want:    pokemonsTest,
			wantErr: nil,
		},
		{
			name: "unsuccess reading pokemons",
			file: nil,
			want: nil,
			wantErr: &model.Error{
				Code:    http.StatusInternalServerError,
				Message: "invalid argument",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErrr := Read(tt.file)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(gotErrr, tt.wantErr) {
				t.Errorf("Read() got1 = %v, want %v", gotErrr, tt.wantErr)
			}
		})
	}
}

func TestReadAllLines(t *testing.T) {
	f, _ := Open("./file/pokemon.csv")
	tests := []struct {
		name    string
		file    *os.File
		want    [][]string
		wantErr *model.Error
	}{
		{
			name:    "success reading all lines",
			file:    f,
			want:    pokemonsTestLines,
			wantErr: nil,
		},
		{
			name: "unsuccess reading all lines",
			file: nil,
			want: nil,
			wantErr: &model.Error{
				Code:    http.StatusInternalServerError,
				Message: "Error trying to read the lines of the file",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErrr := ReadAllLines(tt.file)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadAllLines() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(gotErrr, tt.wantErr) {
				t.Errorf("ReadAllLines() got1 = %v, want %v", gotErrr, tt.wantErr)
			}
		})
	}
}

func TestAddLine(t *testing.T) {
	f, _ := Open("./file/pokemon.csv")
	tests := []struct {
		name     string
		file     *os.File
		lines    [][]string
		newPokes *[]model.SinglePokeExternal
		wantErr  *model.Error
	}{
		{
			name:     "success adding a line",
			file:     f,
			lines:    pokemonsTestLines,
			newPokes: pokemonsFromHttp,
			wantErr:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotErr := AddLine(tt.file, tt.lines, tt.newPokes); !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("AddLine() = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
