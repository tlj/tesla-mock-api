package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tlj/tesla-mock-api/cmd/mock/api"
	"github.com/tlj/tesla-mock-api/cmd/mock/locations"
	"github.com/tlj/tesla-mock-api/cmd/mock/state"
)

func main() {
	state.InitState()

	r := gin.Default()

	r.POST("/oauth/token", api.AuthHandler)

	r.GET("/locations", locations.Locations)

	authorized := r.Group("/api/1", api.AuthRequired)
	authorized.GET("/vehicles", api.Vehicles)

	vehicles := authorized.Group("/vehicles/:vehicle_id", api.VehicleRequiredMiddleware)

	vehicles.GET("/mobile_enabled", api.MobileEnabled)
	vehicles.GET("/data", api.VehicleData)
	vehicles.GET("/data_request/charge_state", api.ChargeState)
	vehicles.GET("/data_request/climate_state", api.ClimateState)
	vehicles.GET("/data_request/drive_state", api.DriveState)
	vehicles.GET("/data_request/gui_settings", api.GuiSettings)
	vehicles.GET("/data_request/vehicle_state", api.VehicleState)
	vehicles.GET("/data_request/vehicle_config", api.VehicleConfig)

	vehicles.POST("/wake_up", api.WakeUp)
	vehicles.POST("/command/auto_conditioning_start", api.AutoConditioningStart)
	vehicles.POST("/command/auto_conditioning_stop", api.AutoConditioningStop)
	vehicles.POST("/command/door_lock", api.DoorLock)
	vehicles.POST("/command/door_unlock", api.DoorUnlock)
	vehicles.POST("/command/remote_seater_request", api.RemoteSeaterRequest)
	vehicles.POST("/command/remote_steering_wheel_heater_request", api.RemoteSteeringWheelHeaterRequest)
	vehicles.POST("/command/set_preconditioning_max", api.SetPreconditioningMax)
	vehicles.POST("/command/set_sentry_mode", api.SetSentryMode)
	vehicles.POST("/command/set_temps", api.SetTemps)

	r.Run()
}
