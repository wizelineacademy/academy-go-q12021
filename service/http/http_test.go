package httpservice

import (
	"pokeapi/model"
	"reflect"
	"testing"
)

func Test_httpService_GetPokemons(t *testing.T) {
	tests := []struct {
		name    string
		h       *httpService
		want    []model.SinglePokeExternal
		wantErr *model.Error
	}{
		{
			name: "succeded pokemon retrieved",
			want: []model.SinglePokeExternal{
				{Name: "delcatty", URL: "https://pokeapi.co/api/v2/pokemon/301/"},
				{Name: "sableye", URL: "https://pokeapi.co/api/v2/pokemon/302/"},
				{Name: "mawile", URL: "https://pokeapi.co/api/v2/pokemon/303/"},
				{Name: "aron", URL: "https://pokeapi.co/api/v2/pokemon/304/"},
				{Name: "lairon", URL: "https://pokeapi.co/api/v2/pokemon/305/"},
			},
			wantErr: nil,
		},
		{
			name: "succeded pokemon retrieved",
			want: []model.SinglePokeExternal{
				{Name: "delcatty", URL: "https://pokeapi.co/api/v2/pokemon/301/"},
				{Name: "sableye", URL: "https://pokeapi.co/api/v2/pokemon/302/"},
				{Name: "mawile", URL: "https://pokeapi.co/api/v2/pokemon/303/"},
				{Name: "aron", URL: "https://pokeapi.co/api/v2/pokemon/304/"},
				{Name: "lairon", URL: "https://pokeapi.co/api/v2/pokemon/305/"},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &httpService{}
			got, gotErr := h.GetPokemons()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpService.GetPokemons() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("httpService.GetPokemons() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
