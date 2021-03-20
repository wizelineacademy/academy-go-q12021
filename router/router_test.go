package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/halarcon-wizeline/academy-go-q12021/router/mocks"
)

//mockgen -source=router/router.go -destination=router/mocks/router.go -package=mocks

func Test_New(t *testing.T) {
	testCases := []struct {
		name           string
		endpoint       string
		handlerName    string
		status         int
		callController bool
	}{
		{
			name:           "OK, Get pokemons",
			endpoint:       "/pokemons",
			handlerName:    "GetLocalPokemons",
			status:         200,
			callController: true,
		},
		{
			name:           "OK, Get pokemon",
			endpoint:       "/pokemons/1",
			handlerName:    "GetLocalPokemon",
			status:         200,
			callController: true,
		},
		{
			name:           "OK, Pokemon api",
			endpoint:       "/pokemons_api",
			handlerName:    "GetExternalPokemons",
			status:         200,
			callController: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			c := mocks.NewMockController(mockCtrl)

			if tc.callController {
				switch tc.handlerName {
					case "GetLocalPokemon":
						c.EXPECT().GetLocalPokemon(gomock.Any(), gomock.Any()).Times(1)
					case "GetLocalPokemons":
						c.EXPECT().GetLocalPokemons(gomock.Any(), gomock.Any()).Times(1)
					case "GetExternalPokemons":
						c.EXPECT().GetExternalPokemons(gomock.Any(), gomock.Any()).Times(1)
				}
			}

			r := New(c)

			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, tc.endpoint, nil)

			r.ServeHTTP(recorder, request)
			assert.Equal(t, tc.status, recorder.Code)
			assert.Nil(t, err)
		})
	}
}
