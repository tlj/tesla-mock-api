package locations

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/tlj/tesla"
	"io/ioutil"
	"net/http"
	"os"
)

func Locations(c *gin.Context) {
	filename := "cmd/mock/fixtures/locations.json"
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatal().Msgf("Error opening %s: %s", filename, err.Error())
	}
	defer jsonFile.Close()

	var locations []tesla.Location
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &locations)
	if err != nil {
		log.Fatal().Msgf("Error parsing %s: %s", filename, err.Error())
	}

	c.JSON(http.StatusOK, locations)
}