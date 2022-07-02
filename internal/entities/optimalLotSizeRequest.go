package entities

type OptimalLotSizeRequest struct {
	RequestId     string    `json:"request_id"`
	Vehicle       []Vehicle `json:"vehicle"`
	NumberOfItems int       `json:"number_of_items"`
	Items         []Item
}

type Item struct {
	Name          string `json:"name"`
	Weight        int    `json:"weight"`
	BoxLengthInCm int    `json:"box_length_in_cm"`
	BoxWidthInCm  int    `json:"box_width_in_cm"`
	BoxHeightInCm int    `json:"box_height_in_cm"`
	ShelfLifeDays bool   `json:"shelf_life_days"`
	Cost          int    `json:"cost"`
}

type Vehicle struct {
	MaxLoadCapacity int `json:"max_load_capacity"`
	HeightInCm      int `json:"height_in_cm"`
	WeightInCm      int `json:"weight_in_cm"`
	LengthInCm      int `json:"length_in_cm"`
}
