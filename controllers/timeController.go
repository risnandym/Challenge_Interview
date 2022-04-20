package controllers

import (
	"interview/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TimeInput struct {
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
	Seconds int `json:"seconds"`
}

const (
	earthtimerule   int = 60
	roketintimerule int = 100
)

// TimeCalculation godoc
// @Summary Time Calculation.
// @Description calculation time from "60" to "100" format.
// @Tags Challenge-1 Time Calculation
// @Param Body body TimeInput true "the body to input earth time."
// @Produce json
// @Router /challenge-1 [post]
func TimeCalculation(ctx *gin.Context) {
	var earth models.Time
	var roketin models.Time
	var input TimeInput
	ctx.BindJSON(&input)

	earth.Hours = input.Hours
	earth.Minutes = input.Minutes
	earth.Seconds = input.Seconds

	// earth hours -->  minutes --> seconds
	earthjam := input.Hours * earthtimerule
	earthmenit := (earthjam + input.Minutes) * earthtimerule
	earthdetik := earthmenit + input.Seconds

	// roketin second --> minutes --> hour
	roketin.Hours = earthdetik / (roketintimerule * roketintimerule)
	roketin.Minutes = (earthdetik % (roketintimerule * roketintimerule)) / roketintimerule
	roketin.Seconds = (earthdetik % (roketintimerule * roketintimerule)) % roketintimerule
	ctx.JSON(http.StatusOK, gin.H{"earth": earth, "roketin": roketin})
}
