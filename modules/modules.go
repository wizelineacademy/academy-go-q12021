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