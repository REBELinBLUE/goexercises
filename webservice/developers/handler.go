package developers

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/REBELinBLUE/goexercises/webservice/response"
	"github.com/julienschmidt/httprouter"
)

// ShowDevelopers retrieves a list of all registered developers
func ShowDevelopers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response.Send(http.StatusOK, Developers, w)
}

// ShowDeveloper retrieves a specific developer
func ShowDeveloper(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("developer"))

	if developer, ok := Developers[int(id)]; ok {
		response.Send(http.StatusOK, developer, w)
		return
	}

	response.Send(http.StatusNotFound, "The developer does not exist", w)
}

// AddDeveloper registers a new developer
func AddDeveloper(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	name := r.FormValue("name")
	age, _ := strconv.Atoi(r.FormValue("age"))
	language := r.FormValue("language")
	floor, _ := strconv.Atoi(r.FormValue("floor"))

	error := ""
	if len(name) == 0 {
		error = "No name supplied"
	} else if len(language) == 0 {
		error = "No language supplied"
	} else if age == 0 {
		error = "No age supplied"
	} else if floor > 5 || floor < -2 {
		error = "Floor must be between -2 and 5"
	}

	if len(error) > 0 {
		response.Send(http.StatusUnprocessableEntity, error, w)
		return
	}

	id := len(Developers) + 1

	developer := Developer{
		Name:     name,
		Age:      age,
		Language: language,
		Floor:    floor,
	}

	Developers[id] = developer

	w.Header().Set("Link", fmt.Sprintf("http://%v/developers/%v", r.Host, id))
	response.Send(http.StatusCreated, developer, w)
}
