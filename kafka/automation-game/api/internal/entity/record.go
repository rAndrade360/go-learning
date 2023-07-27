package entity

import "time"

type Record struct {
	CreatedAt time.Time `json:"createdAt"`
	Sensor    Sensor    `json:"sensor"`
}

type Sensor struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	SerialNumber string `json:"serialNumber"`
	Data         Data   `json:"data"`
}

type Data struct {
	Value string `json:"value"`
}
