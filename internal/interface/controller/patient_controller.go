package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/peetwerapat/hospital-system-api/internal/domain"
	"github.com/peetwerapat/hospital-system-api/internal/interface/controller/dto"
	"github.com/peetwerapat/hospital-system-api/internal/usecase"
	"github.com/peetwerapat/hospital-system-api/pkg/response"
)

type PatientController struct {
	patientUC *usecase.PatientUsecase
}

func NewPatientController(patientUC *usecase.PatientUsecase) *PatientController {
	return &PatientController{patientUC}
}

// @Summary Get Patients
// @Description Search patients by optional filters. Requires staff login.
// @Tags Patient
// @Accept json
// @Produce json
// @Param id query string false "National ID or Passport ID"
// @Param firstName query string false "First Name EN or TH"
// @Param middleName query string false "Middle Name EN or TH"
// @Param lastName query string false "Last Name EN or TH"
// @Param dateOfBirth query string false "Date of Birth (YYYY-MM-DD)"
// @Param phoneNumber query string false "Phone Number"
// @Param email query string false "Email"
// @Success 200 				{object} response.BaseHttpResponse
// @Failure 400,401,500 {object} response.BaseHttpResponse
// @Security ApiKeyAuth
// @Router /patient/search [get]
func (ctrl *PatientController) GetPatientsByHospitalID(c *gin.Context) {
	hospitalIDStr := c.GetString("hospital_id")
	hospitalID, err := strconv.Atoi(hospitalIDStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.BaseHttpResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    usecase.ErrInvalidHospitalToken.Error(),
		})
		c.Abort()
		return
	}

	var filter dto.PatientFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseHttpResponse{
			StatusCode: http.StatusBadRequest,
			Message:    usecase.ErrInvalidQuery.Error(),
		})
		return
	}

	if err := filter.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseHttpResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	filters := map[string]string{
		"id":            filter.ID,
		"first_name":    filter.FirstName,
		"middle_name":   filter.MiddleName,
		"last_name":     filter.LastName,
		"date_of_birth": filter.DateOfBirth,
		"phone_number":  filter.PhoneNumber,
		"email":         filter.Email,
	}

	patients, err := ctrl.patientUC.GetPatientsByHospitalID(hospitalID, filters)
	if err != nil {
		status, msg := usecase.MapGetPatientsByHospitalIDError(err)
		c.JSON(status, response.BaseHttpResponse{
			StatusCode: status,
			Message:    msg,
		})
		return
	}

	c.JSON(http.StatusOK, response.HttpResponseData[[]domain.Patient]{
		BaseHttpResponse: response.BaseHttpResponse{
			StatusCode: http.StatusOK,
			Message:    "Get patients successful",
		},
		Data: patients,
	})
}
