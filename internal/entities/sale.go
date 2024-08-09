package entities

type Sale struct {
	SellingCompany     string `json:"sellingCompany"`
	Product            string `json:"product"`
	Price              string `json:"price"`
	DeliveryCompany    string `json:"deliveryCompany"`
	OriginAddress      string `json:"originAddress"`
	DestinationAddress string `json:"destinationAddress"`
	ClientName         string `json:"clientName"`
	Order              int64  `json:"order"`
}
