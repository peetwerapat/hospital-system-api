package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peetwerapat/hospital-system-api/internal/domain"
	"github.com/peetwerapat/hospital-system-api/internal/interface/controller/dto"
	"github.com/peetwerapat/hospital-system-api/internal/usecase"
	"github.com/peetwerapat/hospital-system-api/pkg/response"
)

type StaffController struct {
	staffUC *usecase.StaffUsecase
}

func NewStaffController(staffUC *usecase.StaffUsecase) *StaffController {
	return &StaffController{staffUC}
}

// @Summary Create staff
// @Description Create staff
// @Tags Staff
// @Accept json
// @Produce json
// @Param staff body dto.StaffRequest true "Staff data"
// @Success 201         	{object} response.BaseHttpResponse
// @Failure 400,404,409,500 {object} response.BaseHttpResponse
// @Router /staff/create [post]
func (ctrl *StaffController) CreateStaff(c *gin.Context) {
	var req dto.StaffRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseHttpResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseHttpResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	staff := &domain.Staff{
		Username:   req.Username,
		Password:   req.Password,
		HospitalID: req.HospitalID,
	}

	if err := ctrl.staffUC.CreateStaff(staff); err != nil {
		status, msg := usecase.MapCreateStaffError(err)
		c.JSON(status, response.BaseHttpResponse{
			StatusCode: status,
			Message:    msg,
		})
		return
	}

	c.JSON(http.StatusCreated, response.BaseHttpResponse{
		StatusCode: http.StatusCreated,
		Message:    "Create staff successful",
	})
}

// @Summary Staff Login
// @Description Staff Login
// @Tags Staff
// @Accept json
// @Produce json
// @Param staff body dto.StaffRequest true "Staff data"
// @Success 200         {object} response.BaseHttpResponse
// @Failure 400,404,500 {object} response.BaseHttpResponse
// @Router /staff/login [post]
func (ctrl *StaffController) StaffLogin(c *gin.Context) {
	var req dto.StaffRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseHttpResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseHttpResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	staff := &domain.Staff{
		Username:   req.Username,
		Password:   req.Password,
		HospitalID: req.HospitalID,
	}

	accessToken, refreshToken, err := ctrl.staffUC.StaffLogin(staff)
	if err != nil {
		status, msg := usecase.MapStaffLoginError(err)
		c.JSON(status, response.BaseHttpResponse{
			StatusCode: status,
			Message:    msg,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"statusCode":   http.StatusOK,
		"message":      "Login successful",
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
