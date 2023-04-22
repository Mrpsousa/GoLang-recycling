package main

import (
	"encoding/json"
	"os"
	"strconv"
)

type SaveRequest struct {
	MSG RequestMSG `json:"msg"`
}

type RequestMSG struct {
	Tp      string  `json:"type"`
	Time    string  `json:"time"`
	Process string  `json:"process"`
	Payload Payload `json:"payload"`
}

type Payload struct {
	Target      Target   `json:"target"`
	WorkorderID string   `json:"workorder_id"`
	Assignees   []string `json:"assignees"`
	Input       Input    `json:"input"`
}

func (payload Payload) UserID() string {
	if len(payload.Assignees) > 0 {
		return payload.Assignees[0]
	}
	return ""
}

type Target struct {
	Tp string `json:"type"`
}

type Input struct {
	Site     string   `json:"site"`
	Facility Facility `json:"facility"`
	Vehicle  Vehicle  `json:"vehicle"`
	Dock     Dock     `json:"dock"`
	Carrier  Carrier  `json:"carrier"`
}

type Facility struct {
	ID string `json:"id"`
}

type Vehicle struct {
	LicensePlate string `json:"license_plate"`
}

type Dock struct {
	name string `json:"name"`
}

func (dock Dock) ID() int {
	i, _ := strconv.Atoi(dock.name)
	return i
}

type Carrier struct {
	id string `json:"id"`
}

func (carrier Carrier) ID() int {
	i, _ := strconv.Atoi(carrier.id)
	return i
}

func main() {
	// conta := Conta{Numero: 1, Saldo: 100}
	// res, err := json.Marshal(conta) //return json in bytes
	// if err != nil {
	// 	println(err)
	// }
	// println(string(res))

	// err = json.NewEncoder(os.Stdout).Encode(conta) // return json "transformed" into json value / serialization process
	// if err != nil {
	// 	println(err)
	// }

	jsonEmpty := []byte(`{
    "type": "WorkOrderCreated",
    "time": "2022-11-16T16:05:02.551692006-04:00",
    "process": "receiving-forward-lego",
    "payload": {
        "workorder_id": "wo_reception_id",
        "target": {
            "type": "package"
        },

        "assignees": ["rep1"],
        "input": {
                "site": "MLB",
                "facility":{
                            "id":"ARENA"
                }
        }
    }
	}`)
	request := SaveRequest{}
	err := json.Unmarshal(jsonEmpty, &request) // json(bytes) to struct
	if err != nil {
		println(err)
	}

	// println(contaX.Saldo)
	json.NewEncoder(os.Stdout).Encode(request)
}
