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

func sendData(w http.ResponseWriter, btc []models.Bitcoin, success bool, errors []string, status int) http.ResponseWriter {
	var data = BitcoinData{}
	data.Data = btc
	data.Success = success
	data.Errors = errors
	json, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(json)
	return w
}

// CreateBitcoin -> Manually does a entry in the db
func CreateBitcoin(w http.ResponseWriter, req *http.Request) {
	var data = BitcoinData{Errors: make([]string, 0)}
	bodyBitcoin, err := helpers.DecodeBody(req)
	if err != nil {
		data.Errors = append(data.Errors, "Could not decode body")
		sendData(w, data.Data, false, data.Errors, http.StatusBadRequest)
		return
	}
	if !helpers.IsValidBase(bodyBitcoin.Base) {
		data.Errors = append(data.Errors, "Invalid Base")
		sendData(w, data.Data, false, data.Errors, http.StatusBadRequest)
		return
	}

	btc, err := models.InsertBitcoinValue(bodyBitcoin.Base, bodyBitcoin.Currency, bodyBitcoin.Amount)
	if err != nil {
		data.Errors = append(data.Errors, "Could not create Entry")
		sendData(w, data.Data, false, data.Errors, http.StatusBadRequest)
		return
	}
	data.Data = append(data.Data, btc)
	sendData(w, data.Data, true, data.Errors, http.StatusOK)
}

// GetBicoins -> gets a list of al bitcoins
func GetBicoins(w http.ResponseWriter, req *http.Request) {
	var btcs []models.Bitcoin = models.GetAllValues()
	var data = BitcoinData{true, btcs, make([]string, 0)}
	sendData(w, data.Data, true, data.Errors, http.StatusOK)
}

// GetBicoin -> gets a list of al bitcoins
func GetBicoin(w http.ResponseWriter, req *http.Request) {
	var data = BitcoinData{}
	vars := mux.Vars(req)

	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		data.Errors = append(data.Errors, "Invalid ID")
		sendData(w, data.Data, false, data.Errors, http.StatusBadRequest)
		return
	}
	btc, err := models.GetValue(ID)
	if err != nil {
		data.Errors = append(data.Errors, "Could not retrive information from db")
		sendData(w, data.Data, false, data.Errors, http.StatusBadRequest)
		return
	}

	data.Data = append(data.Data, btc)

	sendData(w, data.Data, true, nil, http.StatusOK)
}

// GetBitcoinFromAPI -> gets the bitcoin data from coinbase
func GetBitcoinFromAPI(w http.ResponseWriter, r *http.Request) {
	var btc ResponseData
	var data = BitcoinData{}
	response, err := http.Get("https://api.coinbase.com/v2/prices/spot?currency=MXN")
	if err != nil {
		data.Errors = append(data.Errors, "Could not retrive information from coinbase")
		sendData(w, data.Data, false, data.Errors, http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		data.Errors = append(data.Errors, "Could not read from data")
		sendData(w, data.Data, false, data.Errors, http.StatusBadRequest)
		return
	}

	fmt.Println(string(body))

	err = json.Unmarshal(body, &btc)

	if err != nil {
		data.Errors = append(data.Errors, "error decoding data")
		sendData(w, data.Data, false, data.Errors, http.StatusBadRequest)
	}

	btc.Data.ID = 1

	data.Data = append(data.Data, btc.Data)
	sendData(w, data.Data, true, nil, http.StatusOK)
}

// CreateBitcoinFromAPI -> Creates a bitcoin entry based on the coinbase response
func CreateBitcoinFromAPI(w http.ResponseWriter, req *http.Request) {

	var bitcoinReq ResponseData
	var data = BitcoinData{Errors: make([]string, 0)}
	response, err := http.Get("https://api.coinbase.com/v2/prices/spot?currency=MXN")
	if err != nil {
		data.Errors = append(data.Errors, "Could not retrive information from coinbase")
		sendData(w, data.Data, false, data.Errors, http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		data.Errors = append(data.Errors, "Could not read from data")
		sendData(w, data.Data, false, data.Errors, http.StatusBadRequest)
		return
	}

	fmt.Println(string(body))

	err = json.Unmarshal(body, &bitcoinReq)

	if err != nil {
		data.Errors = append(data.Errors, "error decoding data")
		sendData(w, data.Data, false, data.Errors, http.StatusBadRequest)
		return
	}

	btc, err := models.InsertBitcoinValue(bitcoinReq.Data.Base, bitcoinReq.Data.Currency, bitcoinReq.Data.Amount)

	if err != nil {
		data.Errors = append(data.Errors, "could not create Entry")
		sendData(w, data.Data, false, data.Errors, http.StatusBadRequest)
	}

	data.Data = append(data.Data, btc)
	sendData(w, data.Data, true, nil, http.StatusOK)
}

// DeleteBitcoin -> Deletes entry based on id
func DeleteBitcoin(w http.ResponseWriter, req *http.Request) {
	var btcs []models.Bitcoin
	var data = BitcoinData{true, btcs, make([]string, 0)}
	vars := mux.Vars(req)

	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		data.Errors = append(data.Errors, "Invalid ID")
		sendData(w, data.Data, false, data.Errors, http.StatusBadRequest)
		return
	}
	btc, err := models.DeleteValue(ID)
	if err != nil {
		data.Errors = append(data.Errors, "Could not delete entry from db")
		sendData(w, data.Data, false, data.Errors, http.StatusBadRequest)
		return
	}

	data.Data = append(data.Data, btc)
	sendData(w, data.Data, true, nil, http.StatusOK)
}
