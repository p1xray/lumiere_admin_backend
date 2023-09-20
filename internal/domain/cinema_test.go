package domain_test

import (
	"testing"

	"github.com/p1xray/lumiere_admin_backend/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewCinema(t *testing.T) {
	t.Parallel()

	cinemaId := int64(1)
	cinemaName := "test cinema name"
	cinemaDescription := "test cinema description"
	cinemaAddress := "test cinema address"

	testCases := []struct {
		name    string
		cinema  func() (*domain.Cinema, error)
		isValid bool
	}{
		{
			name: "valid",
			cinema: func() (*domain.Cinema, error) {
				return domain.NewCinema(cinemaId, cinemaName, cinemaDescription, cinemaAddress)
			},
			isValid: true,
		},
		{
			name: "invalid cinema id",
			cinema: func() (*domain.Cinema, error) {
				return domain.NewCinema(0, cinemaName, cinemaDescription, cinemaAddress)
			},
			isValid: false,
		},
		{
			name: "missing cinema name",
			cinema: func() (*domain.Cinema, error) {
				return domain.NewCinema(cinemaId, "", cinemaDescription, cinemaAddress)
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cinema, err := tc.cinema()
			if tc.isValid {
				assert.NoError(t, err)

				assert.Equal(t, cinemaId, cinema.Id())
				assert.Equal(t, cinemaName, cinema.Name())
				assert.Equal(t, cinemaDescription, cinema.Description())
				assert.Equal(t, cinemaAddress, cinema.Address())
			} else {
				assert.Error(t, err)
			}
		})
	}
}
