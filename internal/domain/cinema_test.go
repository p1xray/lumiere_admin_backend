package domain_test

import (
	"testing"
	"time"

	"github.com/p1xray/lumiere_admin_backend/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewCinema(t *testing.T) {
	t.Parallel()

	cinemaId := int64(1)
	cinemaName := "test cinema name"
	cinemaDescription := "test cinema description"
	cinemaAddress := "test cinema address"
	cinemaCreatedAt := time.Now()
	cinemaUpdatedAt := time.Now()

	testCases := []struct {
		name    string
		cinema  func() (*domain.Cinema, error)
		isValid bool
	}{
		{
			name: "valid",
			cinema: func() (*domain.Cinema, error) {
				return domain.NewCinema(cinemaId, cinemaName, cinemaDescription, cinemaAddress, cinemaCreatedAt, cinemaUpdatedAt)
			},
			isValid: true,
		},
		{
			name: "invalid cinema id",
			cinema: func() (*domain.Cinema, error) {
				return domain.NewCinema(0, cinemaName, cinemaDescription, cinemaAddress, cinemaCreatedAt, cinemaUpdatedAt)
			},
			isValid: false,
		},
		{
			name: "missing cinema name",
			cinema: func() (*domain.Cinema, error) {
				return domain.NewCinema(cinemaId, "", cinemaDescription, cinemaAddress, cinemaCreatedAt, cinemaUpdatedAt)
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
				assert.Equal(t, cinemaCreatedAt, cinema.CreatedAt())
				assert.Equal(t, cinemaUpdatedAt, cinema.UpdatedAt())
			} else {
				assert.Error(t, err)
			}
		})
	}
}
