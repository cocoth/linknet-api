package request

type ISmartRequest struct {
	FiberNode  string `json:"fiber_node"`
	Address    string `json:"address"`
	Coordinate string `json:"coordinate"`
	Street     string `json:"street"`
}
