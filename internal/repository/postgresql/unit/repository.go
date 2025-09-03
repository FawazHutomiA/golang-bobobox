package unit

import (
	"bobobox/internal/entity"
	"bobobox/pkg/app"
	"bobobox/pkg/helper"
	"bobobox/pkg/sqlx"
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type UnitRepository interface {
	UnitFindAll(ctx context.Context, params helper.PaginationParams, unitType, unitStatus string) (resp helper.Pagination, err error)
	UnitFindByID(ctx context.Context, id uuid.UUID) (resp entity.Unit, err error)
	UnitFindByName(ctx context.Context, name string) (resp entity.Unit, err error)
	UnitInsert(ctx context.Context, unit entity.Unit) (err error)
	UnitUpdateByID(ctx context.Context, unit entity.Unit) (err error)
}

type unitRepository struct {
	app app.AppConfig
}

func NewUnitRepository(app app.AppConfig) UnitRepository {
	return &unitRepository{
		app: app,
	}
}

func (repo *unitRepository) UnitFindAll(ctx context.Context, params helper.PaginationParams, unitType, unitStatus string) (resp helper.Pagination, err error) {
	query := FIND_ALL

	if params.Search != "" {
		escapedSearch := strings.Replace(params.Search, "'", "''", -1)
		addFilter := fmt.Sprintf("AND u.name ILIKE '%%%s%%'", escapedSearch)
		query = fmt.Sprintf(`%s %s`, query, addFilter)
	}

	if unitType != "" {
		addFilter := fmt.Sprintf("AND u.type = '%s'", unitType)
		query = fmt.Sprintf(`%s %s`, query, addFilter)
	}

	if unitStatus != "" {
		addFilter := fmt.Sprintf("AND u.status = '%s'", unitStatus)
		query = fmt.Sprintf(`%s %s`, query, addFilter)
	}

	var unit []entity.Unit

	pagination := sqlx.NewPaginationMetadata(repo.app.Db)
	result, err := pagination.GetPagination(query, params, &unit)
	if err != nil {
		repo.app.Logger.Error(err)
		return resp, err
	}
	return result, nil
}

func (repo *unitRepository) UnitFindByID(ctx context.Context, id uuid.UUID) (resp entity.Unit, err error) {
	err = repo.app.Db.Get(&resp, FIND_BY_ID, id)
	if err != nil {
		repo.app.Logger.Error(err)
		return resp, err
	}
	return resp, nil
}

func (repo *unitRepository) UnitFindByName(ctx context.Context, name string) (resp entity.Unit, err error) {
	err = repo.app.Db.GetContext(ctx, &resp, FIND_BY_NAME, name)
	if err != nil {
		repo.app.Logger.Error(err)
		return resp, err
	}
	return resp, nil
}

func (repo *unitRepository) UnitInsert(ctx context.Context, unit entity.Unit) (err error) {
	_, err = repo.app.Db.ExecContext(ctx, INSERT, unit.ToInsert()...)
	if err != nil {
		repo.app.Logger.Error(err)
		return err
	}
	return nil
}

func (repo *unitRepository) UnitUpdateByID(ctx context.Context, unit entity.Unit) (err error) {
	_, err = repo.app.Db.ExecContext(ctx, UPDATE_BY_ID, unit.ToUpdate()...)
	if err != nil {
		repo.app.Logger.Error(err)
		return err
	}
	return nil
}
