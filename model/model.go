package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Snmp struct {
	ID  primitive.ObjectID      `json:"_id,omitempty" bson:"_id,omitempty"`
	Ip 	string  				`json:"ip,omitempty" bson:"ip,omitempty"`
	Interface string 			`json:"interface,omitempty" bson:"interface,omitempty"`
	Link_status []Link_status   `json:"link_status,omitempty" json:"link_status,omitempty"`
}

type Link_status struct {
	Time int  					`json:"time,omitempty"`
	Status int 					`json:"status,omitempty"`
}

type Avg struct {
	Ip 	string  				`json:"ip,omitempty"`
	Interface string 			`json:"interface,omitempty"`
	Avg float32   				`json:"avg,omitempty"`
}

type Down struct {
	Ip 	string  				`json:"ip,omitempty"`
	Interface string 			`json:"interface,omitempty"`
	Downtime int   				`json:"downtime,omitempty"`
	Countdown int   			`json:"countdown,omitempty"`
	Mttr float32 				`json:"mttr,omitempty"`
}