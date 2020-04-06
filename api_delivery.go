package sdk

import "fmt"

// Delivery :
type Delivery struct {
	DeliveryID       string `json:"deliveryId"`
	DeliveryVendorID string `json:"deliveryVendorId"`
	CurrencyType     string `json:"currencyType"`
	Amount           uint64 `json:"amount"`
	Status           string `json:"status"`
}

// DeliveryType :
type DeliveryType string

const (
	DeliveryTypeDocument DeliveryType = "DOCUMENT"
	DeliveryTypeFood     DeliveryType = "FOOD"
)

// VehicleType :
type VehicleType string

const (
	VehicleTypeMotobike VehicleType = "MOTOBIKE"
	VehicleTypeCar      VehicleType = "CAR"
)

// VendorType :
type VendorType string

const (
	VendorTypeMySpeedy VendorType = "MYSPEEDY"
)

// DeliveryVendor :
type DeliveryVendor struct {
	Vendor     VendorType `json:"vendor"`
	Credential string     `json:"credential"`
}

type DeliveryPoint struct {
	Address        string               `json:"address"`
	EntranceNumber string               `json:"entranceNumber"`
	FloorNumber    string               `json:"floorNumber"`
	BuildingNumber string               `json:"buildingNumber"`
	Remark         string               `json:"remark"`
	Contact        DeliveryPointContact `json:"contact"`
}

type DeliveryPointContact struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}

// RequestCreateDelivery :
type RequestCreateDelivery struct {
	DeliveryVendor DeliveryVendor  `json:"deliveryVendor"`
	VehicleType    VehicleType     `json:"vehicleType"`
	Type           DeliveryType    `json:"type"`
	IsCashAccount  bool            `json:"isCashAccount"`
	Points         []DeliveryPoint `json:"points"`
}

// ResponseCreateDelivery :
type ResponseCreateDelivery struct {
	Code string   `json:"code"`
	Item Delivery `json:"item"`
	Err  *Error   `json:"error"`
}

// CreateDelivery :
func (c Client) CreateDelivery(request RequestCreateDelivery) (*ResponseCreateDelivery, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathCreateDelivery.method
	requestURL := c.prepareAPIURL(pathCreateDelivery)

	response := new(ResponseCreateDelivery)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	return response, nil
}

// ResponseGetDeliveryByID :
type ResponseGetDeliveryByID struct {
	Code string   `json:"code"`
	Item Delivery `json:"item"`
	Err  *Error   `json:"error"`
}

// GetDeliveryByID :
func (c Client) GetDeliveryByID(id string) (*ResponseGetDeliveryByID, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathCreateDelivery.method
	requestURL := c.prepareAPIURL(pathCreateDelivery)
	response := new(ResponseGetDeliveryByID)
	if err := c.httpAPI(method, fmt.Sprintf("%s/%s", requestURL, id), nil, response); err != nil {
		return nil, err
	}

	return response, nil
}

// RequestCalculateDeliveryFee :
type RequestCalculateDeliveryFee struct {
	DeliveryVendor DeliveryVendor  `json:"deliveryVendor"`
	VehicleType    VehicleType     `json:"vehicleType"`
	Type           DeliveryType    `json:"type"`
	IsCashAccount  bool            `json:"isCashAccount"`
	Points         []DeliveryPoint `json:"points"`
}

// ResponseCalculateDeliveryFee :
type ResponseCalculateDeliveryFee struct {
	Code string `json:"code"`
	Err  *Error `json:"error"`
	Item struct {
		CurrencyType string `json:"currencyType"`
		Amount       uint64 `json:"amount"`
	} `json:"item"`
}

// CalculateDeliveryFee :
func (c Client) CalculateDeliveryFee(request RequestCalculateDeliveryFee) (*ResponseCalculateDeliveryFee, error) {
	if c.err != nil {
		return nil, c.err
	}

	method := pathCreateDelivery.method
	requestURL := c.prepareAPIURL(pathCreateDelivery)

	response := new(ResponseCalculateDeliveryFee)
	if err := c.httpAPI(method, requestURL, request, response); err != nil {
		return nil, err
	}

	return response, nil
}