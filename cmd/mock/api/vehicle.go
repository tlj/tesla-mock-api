package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tlj/tesla"
	"github.com/tlj/tesla-mock-api/cmd/mock/state"
	"net/http"
)

func Vehicles(c *gin.Context) {
	var vs []tesla.Vehicle
	for _, v := range state.Current.VehiclesData {
		vs = append(vs, v.ToVehicle())
	}

	c.JSON(http.StatusOK, tesla.VehiclesResponse{Response: vs, Count: int64(len(vs))})
}

func MobileEnabled(c *gin.Context) {
	vs, _ := c.Get("vehicle")
	v := vs.(*tesla.VehicleData)

	c.JSON(http.StatusOK, tesla.BoolResponse{Response: v.MobileEnabled})
}

func VehicleData(c *gin.Context) {
	vs, _ := c.Get("vehicle")
	v := vs.(*tesla.VehicleData)

	c.JSON(http.StatusOK, tesla.VehicleDataResponse{Response: *v})
}

func ChargeState(c *gin.Context) {
	vs, _ := c.Get("vehicle")
	v := vs.(*tesla.VehicleData)

	c.JSON(http.StatusOK, tesla.VehicleChargeStateResponse{Response: v.ChargeState})
}

func ClimateState(c *gin.Context) {
	vs, _ := c.Get("vehicle")
	v := vs.(*tesla.VehicleData)

	c.JSON(http.StatusOK, tesla.VehicleClimateStateResponse{Response: v.ClimateState})
}

func DriveState(c *gin.Context) {
	vs, _ := c.Get("vehicle")
	v := vs.(*tesla.VehicleData)

	c.JSON(http.StatusOK, tesla.VehicleDriveStateResponse{Response: v.DriveState})
}

func GuiSettings(c *gin.Context) {
	vs, _ := c.Get("vehicle")
	v := vs.(*tesla.VehicleData)

	c.JSON(http.StatusOK, tesla.VehicleGuiSettingsResponse{Response: v.GuiSettings})
}

func VehicleState(c *gin.Context) {
	vs, _ := c.Get("vehicle")
	v := vs.(*tesla.VehicleData)

	c.JSON(http.StatusOK, tesla.VehicleVehicleStateResponse{Response: v.VehicleState})
}

func VehicleConfig(c *gin.Context) {
	vs, _ := c.Get("vehicle")
	v := vs.(*tesla.VehicleData)

	c.JSON(http.StatusOK, tesla.VehicleVehicleConfigResponse{Response: v.VehicleConfig})
}
