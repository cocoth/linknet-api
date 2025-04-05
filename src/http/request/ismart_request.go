package request

type ISmartRequest struct {
	FiberNode    string `json:"fiber_node"`
	Address      string `json:"address"`
	Coordinate   string `json:"coordinate"`
	CustomerName string `json:"customer_name"`
	Street       string `json:"street"`
}
