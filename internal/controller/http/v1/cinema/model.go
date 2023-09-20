package cinema

// Модель кинотеатра для запросов
type Cinema struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
}
