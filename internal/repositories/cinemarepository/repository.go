package cinemarepository

import "github.com/p1xray/lumiere_admin_backend/pkg/postgres"

type CinemaRepository struct {
	pg *postgres.Postgres
}

func NewCinemaRepository(pg *postgres.Postgres) *CinemaRepository {
	return &CinemaRepository{
		pg: pg,
	}
}
