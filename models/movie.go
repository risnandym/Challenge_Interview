package models

type (
	Movie struct {
		ID          uint   `gorm:"primary_key" json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Duration    int    `json:"duration"`
		Artists     string `json:"artists"`
		Genres      string `json:"genres"`
	}

	Pagination struct {
		Limit int `json:"limit"`
		Page  int `json:"page"`
	}
)
