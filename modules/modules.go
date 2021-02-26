// func UpdatePokemon(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id := params["id"]

// 	if !bson.IsObjectIdHex(id) {
// 		responseWithError(w, http.StatusBadRequest)
// 	}

// 	var poke model.Pokemon
// 	decoder := json.NewDecoder(r.Body)
// 	objectId := bson.ObjectIdHex(id)

// 	err := decoder.Decode(&poke)

// 	if err != nil {
// 		responseWithError(w, http.StatusInternalServerError)
// 	}
	
// 	defer r.Body.Close()

// 	document := bson.M{"_id": objectId}
// 	change := bson.M{"$set":poke}

// 	err = collection.Update(document, change)

// 	if err != nil {
// 		responseWithError(w, http.StatusInternalServerError)
// 	} else {
// 		responseOne(w, poke)
// 	}
// }

// func DeletePokemon(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id := params["id"]

// 	if !bson.IsObjectIdHex(id) {
// 		responseWithError(w, http.StatusBadRequest)
// 	}

// 	var poke model.Pokemon
// 	objectId := bson.ObjectIdHex(id)

// 	err := collection.FindId(objectId).One(&poke)

// 	if err != nil {
// 		responseWithError(w, http.StatusInternalServerError)
// 	}

// 	err = collection.RemoveId(objectId)

// 	if err != nil {
// 		responseWithError(w, http.StatusInternalServerError)
// 	} else {
// 		responseOne(w, poke)
// 	}
// }