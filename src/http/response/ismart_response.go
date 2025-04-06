package response

import "time"

type ISmartResponse struct {
	ID         string     `json:"id"`
	FiberNode  string     `json:"fiber_node"`
	Address    string     `json:"address"`
	Coordinate string     `json:"coordinate"`
	Street     string     `json:"street"`
	CreatedAt  time.Time  `json:"CreatedAt"`
	UpdatedAt  time.Time  `json:"UpdatedAt"`
	DeletedAt  *time.Time `json:"DeletedAt"`
}
