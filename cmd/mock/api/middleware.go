package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tlj/tesla"
	"github.com/tlj/tesla-mock-api/cmd/mock/state"
	"net/http"
)

func VehicleRequiredMiddleware(c *gin.Context) {
	vehicleID, _ := c.Params.Get("vehicle_id")
	v := state.Current.VehiclesData.ByIDs(vehicleID)
	if v == nil {
		c.JSON(http.StatusBadRequest, tesla.BoolResponse{Response: false})
		return
	}

	c.Set("vehicle", v)

	c.Next()
}