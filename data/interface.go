package data

import (
	"time"
)

type Id string

type Dose struct {
	Id        Id
	Time      time.Time
	Quantity  string
	Substance string
	Route     string
}

type JsonDose struct {
	Id        Id
	Time      int64
	Quantity  string
	Substance string
	Route     string
}

type Storage interface {
	FetchAll() ([]Dose, error)
	Append(*Dose) error
	DeleteDose(Id) error
}

var ApplicationStorage Storage

func (this JsonDose) ToDose() Dose {
	return Dose{
		Id:        this.Id,
		Time:      time.Unix(this.Time, 0),
		Quantity:  this.Substance,
		Substance: this.Substance,
		Route:     this.Route,
	}
}

func (this Dose) ToJsonDose() JsonDose {
	return JsonDose{
		Id:        this.Id,
		Time:      this.Time.Unix(),
		Quantity:  this.Substance,
		Substance: this.Substance,
		Route:     this.Route,
	}
}
