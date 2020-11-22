package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tlj/tesla"
	"github.com/tlj/tesla-mock-api/cmd/mock/state"
	"net/http"
)

func SetPreconditioningMax(c *gin.Context) {
	var req struct {
		On bool `form:"on"`
	}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, tesla.BoolReasonResponse{Response: false, Reason: "missing or invalid 'on' param"})
		return
	}

	vehicleID, _ := c.Params.Get("vehicle_id")
	v := state.Current.VehiclesData.ByIDs(vehicleID)
	v.ClimateState.IsFrontDefrosterOn = req.On
	v.ClimateState.IsRearDefrosterOn = req.On
	v.ClimateState.IsPreconditioning = req.On

	c.JSON(http.StatusOK, tesla.BoolReasonResponse{Response: true})
}

func AutoConditioningStart(c *gin.Context) {
	vehicleID, _ := c.Params.Get("vehicle_id")
	v := state.Current.VehiclesData.ByIDs(vehicleID)
	v.ClimateState.IsAutoConditioningOn = true

	c.JSON(http.StatusOK, tesla.BoolReasonResponse{Response: true})
}

func AutoConditioningStop(c *gin.Context) {
	vehicleID, _ := c.Params.Get("vehicle_id")
	v := state.Current.VehiclesData.ByIDs(vehicleID)
	v.ClimateState.IsAutoConditioningOn = false

	c.JSON(http.StatusOK, tesla.BoolReasonResponse{Response: true})
}

func SetTemps(c *gin.Context) {
	var req struct {
		DriverTemp    float64 `form:"driver_temp"`
		PassengerTemp float64 `form:"passenger_temp"`
	}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, tesla.BoolReasonResponse{Response: false, Reason: "missing or invalid params"})
		return
	}

	vehicleID, _ := c.Params.Get("vehicle_id")
	v := state.Current.VehiclesData.ByIDs(vehicleID)
	v.ClimateState.DriverTempSetting = req.DriverTemp
	v.ClimateState.PassengerTempSetting = req.PassengerTemp

	c.JSON(http.StatusOK, tesla.BoolReasonResponse{Response: true})
}

func RemoteSeaterRequest(c *gin.Context) {
	var req struct {
		Heater int64 `form:"heater"`
		Level  int64 `form:"level"`
	}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, tesla.BoolReasonResponse{Response: false, Reason: "missing or invalid params"})
		return
	}

	vehicleID, _ := c.Params.Get("vehicle_id")
	v := state.Current.VehiclesData.ByIDs(vehicleID)
	if req.Heater == 0 {
		v.ClimateState.SeatHeaterLeft = req.Level
	}
	if req.Heater == 1 {
		v.ClimateState.SeatHeaterRight = req.Level
	}

	c.JSON(http.StatusOK, tesla.BoolReasonResponse{Response: true})
}

func RemoteSteeringWheelHeaterRequest(c *gin.Context) {
	var req struct {
		On bool `form:"on"`
	}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, tesla.BoolReasonResponse{Response: false, Reason: "missing or invalid 'on' param"})
		return
	}

	c.JSON(http.StatusOK, tesla.BoolReasonResponse{Response: true})
}
