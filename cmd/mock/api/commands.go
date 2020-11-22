package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tlj/tesla"
	"github.com/tlj/tesla-mock-api/cmd/mock/state"
	"net/http"
	"time"
)

func SetSentryMode(c *gin.Context) {
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
	if v.VehicleState.SentryModeAvailable {
		v.VehicleState.SentryMode = req.On
	} else {
		v.VehicleState.SentryMode = false
	}

	c.JSON(http.StatusOK, tesla.BoolReasonResponse{Response: true})
}

func DoorUnlock(c *gin.Context) {
	vehicleID, _ := c.Params.Get("vehicle_id")
	v := state.Current.VehiclesData.ByIDs(vehicleID)
	v.VehicleState.Locked = false

	c.JSON(http.StatusOK, tesla.BoolReasonResponse{Response: true})
}

func DoorLock(c *gin.Context) {
	vehicleID, _ := c.Params.Get("vehicle_id")
	v := state.Current.VehiclesData.ByIDs(vehicleID)
	v.VehicleState.Locked = true

	c.JSON(http.StatusOK, tesla.BoolReasonResponse{Response: true})
}

func WakeUp(c *gin.Context) {
	vehicleID, _ := c.Params.Get("vehicle_id")
	v := state.Current.VehiclesData.ByIDs(vehicleID)

	if v.State == "asleep" {
		v.State = "waking"

		go func(iv *tesla.VehicleData) {
			timer := time.NewTimer(30 * time.Second)
			for {
				select {
				case <-timer.C:
					v.State = "online"
					timer.Stop()
					break
				}
			}
		}(v)
	}

	c.JSON(http.StatusOK, tesla.VehicleResponse{Response: v.ToVehicle()})
}

