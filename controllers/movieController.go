package controllers

import (
	"interview/config"
	"interview/models"
	"interview/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type movieInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Artists     string `json:"artists"`
	Genres      string `json:"genres"`
}

// CreateMovie godoc
// @Summary Create New Movie.
// @Description Creating a new Movie.
// @Tags Movie
// @Param Body body movieInput true "the body to create a new movie"
// @Produce json
// @Success 200 {object} models.Movie
// @Router /movies [post]
func CreateMovie(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input movieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Movie
	movie := models.Movie{
		Title:       input.Title,
		Description: input.Description,
		Duration:    input.Duration,
		Artists:     input.Artists,
		Genres:      input.Genres,
	}
	db.Create(&movie)

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// UpdateMovie godoc
// @Summary Update Movie.
// @Description Update movie by id.
// @Tags Movie
// @Produce json
// @Param id path string true "movie id"
// @Param Body body movieInput true "the body to update an movie"
// @Success 200 {object} models.Movie
// @Router /movies/{id} [patch]
func UpdateMovie(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var movie models.Movie
	if err := db.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input movieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Movie
	updatedInput.Title = input.Title
	updatedInput.Description = input.Description
	updatedInput.Duration = input.Duration
	updatedInput.Artists = input.Artists
	updatedInput.Genres = input.Genres

	db.Model(&movie).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// DeleteMovie godoc
// @Summary Delete one movie.
// @Description Delete a movie by id.
// @Tags Movie
// @Produce json
// @Param id path string true "movie id"
// @Success 200 {object} map[string]boolean
// @Router /movie/{id} [delete]
func DeleteMovie(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var movie models.Movie
	if err := db.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&movie)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// GetAllMovies godoc
// @Summary Get all movies.
// @Description Get a list of Movies.
// @Tags Movie
// @Produce json
// @Success 200 {object} []models.Movie
// @Router /movies [get]
func GetAllMovie(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	pagination := utils.GeneratePaginationFromRequest(c)
	var movies models.Movie
	movieLists, err := withpagination(&movies, &pagination)
	db.Find(&movies)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{"data": movieLists})
}

// GetMovieById godoc
// @Summary Get Movie.
// @Description Get a Movie by id.
// @Tags Movie
// @Produce json
// @Param id path string true "movie id"
// @Success 200 {object} models.Movie
// @Router /movies/{id} [get]
func GetMovieById(c *gin.Context) { // Get model if exist
	var movie models.Movie

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

func withpagination(movie *models.Movie, pagination *models.Pagination) (*[]models.Movie, error) {
	var movies []models.Movie
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := config.DB.Limit(pagination.Limit).Offset(offset)
	result := queryBuilder.Model(&models.Movie{}).Where(movie).Find(&movies)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &movies, nil
}
