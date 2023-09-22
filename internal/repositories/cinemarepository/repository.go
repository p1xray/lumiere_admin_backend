package cinemarepository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/p1xray/lumiere_admin_backend/pkg/postgres"
)

// Репозиторий кинотеатра
type CinemaRepository struct {
	pg *postgres.Postgres
}

// Возвращает новый репозиторий кинотеатра
func NewCinemaRepository(pg *postgres.Postgres) *CinemaRepository {
	return &CinemaRepository{
		pg: pg,
	}
}

// Возвращает список кинотеатров из хранилища
func (cr *CinemaRepository) GetList(ctx context.Context) ([]Cinema, error) {
	sql, _, err := cr.pg.Builder.
		Select("id, name, description, address, created_at, updated_at").
		From("cinemas").
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := cr.pg.Pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	entities := make([]Cinema, 0)
	for rows.Next() {
		e := Cinema{}
		err := rows.Scan(&e.Id, &e.Name, &e.Description, &e.Address, &e.CreatedAt, &e.UpdatedAt)
		if err != nil {
			return nil, err
		}

		entities = append(entities, e)
	}

	return entities, nil
}

// Возвращает подробности кинотеатра
func (cr *CinemaRepository) GetDetails(ctx context.Context, id int64) (*Cinema, error) {
	sql, args, err := cr.pg.Builder.
		Select("id, name, description, address, created_at, updated_at").
		From("cinemas").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return nil, err
	}

	e := &Cinema{}
	row := cr.pg.Pool.QueryRow(ctx, sql, args...)
	err = row.Scan(&e.Id, &e.Name, &e.Description, &e.Address, &e.CreatedAt, &e.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return e, nil
}
