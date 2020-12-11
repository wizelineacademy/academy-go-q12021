package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"../db/models"
	"../helpers"
	"github.com/gorilla/mux"
)

// BitcoinData -> this structure is the base response for the project for bitcoins
type BitcoinData struct {
	Success bool             `json:"success"`
	Data    []models.Bitcoin `json:"data"`
	Errors  []string         `json:"errors"`
}

// ResponseData -> this structure is used to adapt to the response sent by the coinbase api
type ResponseData struct {
	Data models.Bitcoin `json:"data"`
}

// CreateBitcoin -> Manually does a entry in the db
func CreateBitcoin(w http.ResponseWriter, req *http.Request) {
	bodyBitcoin, success := helpers.DecodeBody(req)
	if success != true {
		http.Error(w, "could not decode body", http.StatusBadRequest)
		return
	}

	var data = BitcoinData{Errors: make([]string, 0)}
	if !helpers.IsValidBase(bodyBitcoin.Base) {
		data.Success = false
		data.Errors = append(data.Errors, "invalid base")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	btc, success := models.InsertBitcoinValue(bodyBitcoin.Base, bodyBitcoin.Currency, bodyBitcoin.Amount)
	if success != true {
		data.Errors = append(data.Errors, "could not create Entry")
	}

	data.Success = success
	data.Data = append(data.Data, btc)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
	return
}

// GetBicoins -> gets a list of al bitcoins
func GetBicoins(w http.ResponseWriter, req *http.Request) {
	var btcs []models.Bitcoin = models.GetAllValues()

	var data = BitcoinData{true, btcs, make([]string, 0)}
	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

// GetBicoin -> gets a list of al bitcoins
func GetBicoin(w http.ResponseWriter, req *http.Request) {
	var btcs []models.Bitcoin
	var data = BitcoinData{true, btcs, make([]string, 0)}
	vars := mux.Vars(req)

	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		data.Success = false
		data.Errors = append(data.Errors, "Invalid ID")
		json, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
		return
	}
	btc, success := models.GetValue(ID)
	if success != true {
		data.Errors = append(data.Errors, "Could not retrive information from db")
		json, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
		return
	}

	data.Data = append(data.Data, btc)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

// GetBitcoinFromAPI -> gets the bitcoin data from coinbase
func GetBitcoinFromAPI(w http.ResponseWriter, r *http.Request) {
	var btc ResponseData
	response, err := http.Get("https://api.coinbase.com/v2/prices/spot?currency=MXN")
	if err != nil {
		fmt.Fprintf(w, "Could not retrive information from coinbase")
		return
	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Fprintf(w, "Could not read from data")
		return
	}

	fmt.Println(string(data))

	err = json.Unmarshal(data, &btc)

	if err != nil {
		fmt.Println(err)
	}

	btc.Data.ID = 1

	fmt.Fprint(w, btc)

}

// CreateBitcoinFromAPI -> Creates a bitcoin entry based on the coinbase response
func CreateBitcoinFromAPI(w http.ResponseWriter, req *http.Request) {

	var bitcoinReq ResponseData
	response, err := http.Get("https://api.coinbase.com/v2/prices/spot?currency=MXN")
	if err != nil {
		fmt.Fprintf(w, "Could not retrive information from coinbase")
		return
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Fprintf(w, "Could not read from data")
		return
	}

	fmt.Println(string(body))

	err = json.Unmarshal(body, &bitcoinReq)

	if err != nil {
		fmt.Println(err)
		return
	}

	btc, success := models.InsertBitcoinValue(bitcoinReq.Data.Base, bitcoinReq.Data.Currency, bitcoinReq.Data.Amount)

	var data = BitcoinData{Errors: make([]string, 0)}

	if success != true {
		data.Errors = append(data.Errors, "could not create Entry")
	}

	data.Success = success
	data.Data = append(data.Data, btc)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
	return
}

// DeleteBitcoin -> Deletes entry based on id
func DeleteBitcoin(w http.ResponseWriter, req *http.Request) {
	var btcs []models.Bitcoin
	var data = BitcoinData{true, btcs, make([]string, 0)}
	vars := mux.Vars(req)

	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		data.Errors = append(data.Errors, "Invalid ID")
		json, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
		return
	}
	btc, success := models.DeleteValue(ID)
	if success != true {
		data.Errors = append(data.Errors, "Could not delete entry from db")
		json, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
		return
	}

	data.Data = append(data.Data, btc)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)

}
