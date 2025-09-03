package unit

import (
	"bobobox/pkg/app"
	"bobobox/pkg/constant"
	"bobobox/pkg/helper"
	"bobobox/pkg/response"
	"bobobox/pkg/validator"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type UnitHandler struct {
	App         app.AppConfig
	UnitService UnitService
}

func NewUnitHandler(app app.AppConfig, unitService UnitService) *UnitHandler {
	return &UnitHandler{App: app, UnitService: unitService}
}

func (handler *UnitHandler) ListPaginate(w http.ResponseWriter, r *http.Request) {
	// Init
	var resp response.Response
	ctx := r.Context()

	param := helper.PaginationParams{}
	param = param.GetPaginateParam(r)

	unitType := r.URL.Query().Get("type")
	unitStatus := r.URL.Query().Get("status")

	service, err := handler.UnitService.ListPaginate(ctx, param, unitType, unitStatus)
	if err.Errors != nil {
		handler.App.Logger.Error(err)
		resp = response.Error(err.Status, err.Message, err.Errors)
		resp.JSON(w)
		return
	}

	resp = response.Success(response.StatusOK, "Success", service)
	resp.JSON(w)
}

func (handler *UnitHandler) Detail(w http.ResponseWriter, r *http.Request) {
	// Init
	var resp response.Response
	ctx := r.Context()

	idStr := chi.URLParam(r, "id")

	id, errs := uuid.Parse(idStr)
	if errs != nil {
		handler.App.Logger.Error(errs)
		resp = response.Error(response.StatusBadRequest, constant.StatusBadRequest, errs)
		resp.JSON(w)
		return
	}

	service, err := handler.UnitService.Detail(ctx, id)
	if err.Errors != nil {
		handler.App.Logger.Error(err)
		resp = response.Error(err.Status, err.Message, err.Errors)
		resp.JSON(w)
		return
	}

	resp = response.Success(response.StatusOK, "Success", service)
	resp.JSON(w)
}

func (handler *UnitHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Init
	var req UnitCreateRequest
	var resp response.Response
	ctx := r.Context()

	resp, errV := validator.ValidateRequest(r, &req)
	if errV != nil {
		resp.JSON(w)
		return
	}

	service, err := handler.UnitService.Create(ctx, req)
	if err.Errors != nil {
		handler.App.Logger.Error(err)
		resp = response.Error(err.Status, err.Message, err.Errors)
		resp.JSON(w)
		return
	}

	resp = response.Success(response.StatusOK, "Success", service)
	resp.JSON(w)
}

func (handler *UnitHandler) Update(w http.ResponseWriter, r *http.Request) {
	// Init
	var req UnitUpdateRequest
	var resp response.Response
	ctx := r.Context()

	idStr := chi.URLParam(r, "id")

	id, errs := uuid.Parse(idStr)
	if errs != nil {
		handler.App.Logger.Error(errs)
		resp = response.Error(response.StatusBadRequest, constant.StatusBadRequest, errs)
		resp.JSON(w)
		return
	}

	resp, errV := validator.ValidateRequest(r, &req)
	if errV != nil {
		resp.JSON(w)
		return
	}

	service, err := handler.UnitService.Update(ctx, id, req)
	if err.Errors != nil {
		handler.App.Logger.Error(err)
		resp = response.Error(err.Status, err.Message, err.Errors)
		resp.JSON(w)
		return
	}

	resp = response.Success(response.StatusOK, "Success", service)
	resp.JSON(w)
}
