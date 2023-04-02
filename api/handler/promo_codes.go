package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Promo Code godoc
// @ID create_promo_code
// @Router /promo_codes [POST]
// @Summary Create Promo Code
// @Description Create Promo Code
// @Tags Promo Code
// @Accept json
// @Produce json
// @Param promo_codes body models.CreatePromo true "CreatePromoCodeRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreatePromoCode(c *gin.Context) {

	var createPromo models.CreatePromo

	err := c.ShouldBindJSON(&createPromo) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create promo code", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.PromoCode().Create(context.Background(), &createPromo)
	if err != nil {
		h.handlerResponse(c, "storage.promo.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.PromoCode().GetByID(context.Background(), &models.PromoPrimaryKey{Promo_id: id})
	if err != nil {
		h.handlerResponse(c, "storage.promo.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create promo", http.StatusCreated, resp)
}

// Get By ID Promo Code godoc
// @ID get_by_id_promo_code
// @Router /promo_codes/{id} [GET]
// @Summary Get By ID Promo Code
// @Description Get By ID Promo Code
// @Tags Promo Code
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdPromoCode(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.promo.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

		resp, err := h.storages.PromoCode().GetByID(context.Background(), &models.PromoPrimaryKey{Promo_id: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.promo.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get promo by id", http.StatusCreated, resp)
}


// Get List Promo Code godoc
// @ID get_list_promo
// @Router /promo_codes [GET]
// @Summary Get List Promo
// @Description Get List Promo
// @Tags Promo Code
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListPromoCode(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list promo", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list promo", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.PromoCode().GetList(context.Background(), &models.GetListPromoRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.promo.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list promo response", http.StatusOK, resp)
}

// DELETE Promo godoc
// @ID delete promo_codes
// @Router /promo_codes/{id} [DELETE]
// @Summary Delete Promo Code
// @Description Delete Promo Code
// @Tags Promo Code
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 204 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeletePromoCode(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.promo_code.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	rowsAffected, err := h.storages.PromoCode().Delete(context.Background(), &models.PromoPrimaryKey{Promo_id: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.promo_code.delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.promo_code.delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "delete promo_code", http.StatusNoContent, nil)
}