package user

import (
	"net/http"

	"github.com/codoworks/go-boilerplate/pkg/api/handlers"
	"github.com/codoworks/go-boilerplate/pkg/api/helpers"
	"github.com/codoworks/go-boilerplate/pkg/db/models"
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	ms, err := models.UserModel().FindAll()

	if err != nil {
		return helpers.Error(c, err, nil)
	}
	var payload []*models.UserForm
	for _, m := range ms {
		f := m.MapToForm()
		payload = append(payload, f)
	}
	return c.JSON(http.StatusOK, handlers.Success(payload))
}
