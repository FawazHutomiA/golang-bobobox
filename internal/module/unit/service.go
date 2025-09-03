package unit

import (
	"context"

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

	return resp, errData
}

// Create usecase
func (uc *unitService) Create(ctx context.Context, params UnitCreateRequest) (resp UnitCreateResponse, errData exception.Error) {

	return resp, errData
}

func (uc *unitService) Update(ctx context.Context, id uuid.UUID, params UnitUpdateRequest) (resp UnitUpdateResponse, errData exception.Error) {

	return resp, errData
}
