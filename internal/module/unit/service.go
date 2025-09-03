package unit

import (
	"context"
	"database/sql"

	"bobobox/internal/entity"
	"bobobox/internal/repository/postgresql/unit"
	"bobobox/pkg/app"
	"bobobox/pkg/exception"
	"bobobox/pkg/helper"
	"bobobox/pkg/response"

	"github.com/google/uuid"
)

type UnitService interface {
	ListPaginate(ctx context.Context, params helper.PaginationParams, unitType, unitStatus string) (resp helper.Pagination, errData exception.Error)
	Detail(ctx context.Context, id uuid.UUID) (resp UnitDetailResponse, errData exception.Error)
	Create(ctx context.Context, params UnitCreateRequest) (resp UnitCreateResponse, errData exception.Error)
	Update(ctx context.Context, id uuid.UUID, params UnitUpdateRequest) (resp UnitUpdateResponse, errData exception.Error)
}

type unitService struct {
	app        app.AppConfig
	repository unit.UnitRepository
}

func NewUnitService(app app.AppConfig, repository unit.UnitRepository) UnitService {
	return &unitService{
		app:        app,
		repository: repository,
	}
}

func (uc *unitService) ListPaginate(ctx context.Context, params helper.PaginationParams, unitType, unitStatus string) (resp helper.Pagination, errData exception.Error) {
	// repository
	unitRepo, err := uc.repository.UnitFindAll(ctx, params, unitType, unitStatus)
	if err != nil {
		return resp, exception.Error{
			Status:  response.StatusBadRequest,
			Message: "Something Wrong",
			Errors:  exception.ErrBadRequest,
		}
	}

	return unitRepo, errData
}

func (uc *unitService) Detail(ctx context.Context, id uuid.UUID) (resp UnitDetailResponse, errData exception.Error) {
	// repository
	unitRepo, err := uc.repository.UnitFindByID(ctx, id)

	switch err {
	case nil:
		err = nil
	case sql.ErrNoRows:
		return resp, exception.Error{
			Status:  response.StatusNotFound,
			Message: "Unit  Not Found",
			Errors:  exception.ErrNotFound,
		}
	default:
		return resp, exception.Error{
			Status:  response.StatusBadRequest,
			Message: "Something Wrong",
			Errors:  exception.ErrBadRequest,
		}
	}

	// map response
	resp = UnitDetailResponse(unitRepo)

	return resp, errData
}

// Create usecase
func (uc *unitService) Create(ctx context.Context, params UnitCreateRequest) (resp UnitCreateResponse, errData exception.Error) {
	// repository
	_, err := uc.repository.UnitFindByName(ctx, params.Name)
	switch err {
	case nil:
		return resp, exception.Error{
			Status:  response.StatusConflicted,
			Message: "This Unit  Already Exsist",
			Errors:  exception.ErrConflicted,
		}
	case sql.ErrNoRows:
		err = nil
	default:
		return resp, exception.Error{
			Status:  response.StatusBadRequest,
			Message: "Something Wrong",
			Errors:  exception.ErrBadRequest,
		}
	}

	unitID := uuid.New()

	// map insert
	unit := entity.Unit{
		ID:     unitID,
		Name:   params.Name,
		Type:   params.Type,
		Status: params.Status,
	}

	// save to db
	err = uc.repository.UnitInsert(ctx, unit)
	if err != nil {
		return resp, exception.Error{
			Status:  response.StatusInternalServerError,
			Message: "Error",
			Errors:  exception.ErrInternalServer,
		}
	}

	params.ID = unitID

	// map response
	resp = UnitCreateResponse(params)

	return resp, errData
}

func (uc *unitService) Update(ctx context.Context, id uuid.UUID, params UnitUpdateRequest) (resp UnitUpdateResponse, errData exception.Error) {

	return resp, errData
}
