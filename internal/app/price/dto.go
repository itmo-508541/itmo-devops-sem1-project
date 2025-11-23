package price

type PriceDTO struct {
	Id         string `csv:"id"          json:"id"          validate:"notblank,number"`
	Name       string `csv:"name"        json:"name"        validate:"notblank"`
	Category   string `csv:"category"    json:"category"    validate:"notblank"`
	Price      string `csv:"price"       json:"price"       validate:"notblank,numeric"`
	CreateDate string `csv:"create_date" json:"create_date" validate:"notblank,date"`
}

type AcceptResultDTO struct {
	TotalCount      int     `json:"total_count"`
	DuplicatesCount int     `json:"duplicates_count"`
	TotalItems      int     `json:"total_items"`
	TotalCategories int     `json:"total_categories"`
	TotalPrice      float32 `json:"total_price"`
}
