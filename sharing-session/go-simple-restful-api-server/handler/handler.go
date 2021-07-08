package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	jsonData := make(map[string]interface{})
	jsonData["message"] = "Welcome to CloudKilat"
	jsonData["status"] = "success"
	jsonData["code"] = http.StatusOK

	fmt.Println("jsonData:", jsonData)

	// encode to json format
	enc, err := json.Marshal(jsonData)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Println(string(enc))
	w.Write(enc)
}

type VMServices struct {
	ID      uint32 `json:"id"`
	Name    string `json:"name"`
	Package string `json:"package"`
	Price   uint32 `json:"price"`

	Specification `json:"specification"`
}

type Specification struct {
	CPU string `json:"cpu"`
	RAM string `json:"ram"`
}

func GetVMServicesHandler(w http.ResponseWriter, r *http.Request) {
	vm := []VMServices{
		{
			ID:      1,
			Name:    "Kilat VM 2.0",
			Package: "XXS",
			Price:   90000,
			Specification: Specification{
				CPU: "1 Core",
				RAM: "2 GB",
			},
		},
		{
			ID:      2,
			Name:    "Kilat VM 2.0",
			Package: "XS",
			Price:   180000,
			Specification: Specification{
				CPU: "2 Core",
				RAM: "4 gb",
			},
		},
	}

	enc, err := json.Marshal(vm)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(enc)
}
