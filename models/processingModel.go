package model

type Processing struct {
	Id         int     `json:"id"`
	Process    float64 `json:"processing" binding:"required"`
	FreeMemory float64 `json:"free_memory" binding:"required"`
	Swap       float64 `json:"swap" binding:"required"`
}
