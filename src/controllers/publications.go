package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userIDToken, err := auth.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication models.Publications
	if err = json.Unmarshal(requestBody, &publication); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	publication.AuthorID = userIDToken

	if err = publication.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationsRepository(db)
	publication.ID, err = repository.Create(publication)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, publication)
}

func GetPublications(w http.ResponseWriter, r *http.Request) {
	userIDToken, err := auth.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationsRepository(db)
	publications, err := repository.Search(userIDToken)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publications)
}

func GetPublication(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	publicationID, err := strconv.ParseUint(parameters["publicationID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationsRepository(db)
	publication, err := repository.SearchByID(publicationID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publication)
}

func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	userIDToken, err := auth.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	publicationID, err := strconv.ParseUint(parameters["publicationID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationsRepository(db)

	databasePublication, err := repository.SearchByID(publicationID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if databasePublication.AuthorID != userIDToken {
		responses.Error(w, http.StatusForbidden, errors.New("não é possível atualizar uma publicação que não seja sua"))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication models.Publications
	if err = json.Unmarshal(requestBody, &publication); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = publication.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.Update(publicationID, publication); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

func DeletePublication(w http.ResponseWriter, r *http.Request) {
	userIDToken, err := auth.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	publicationID, err := strconv.ParseUint(parameters["publicationID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationsRepository(db)

	databasePublication, err := repository.SearchByID(publicationID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if databasePublication.AuthorID != userIDToken {
		responses.Error(w, http.StatusForbidden, errors.New("não é possível deletar uma publicação que não seja sua"))
		return
	}

	if err = repository.Delete(publicationID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func GetPublicationByUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationsRepository(db)

	publications, err := repository.SearchByUser(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publications)
}

func LikePublication(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	publicationID, err := strconv.ParseUint(parameters["publicationID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationsRepository(db)

	if err = repository.Like(publicationID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func UnlikePublication(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	publicationID, err := strconv.ParseUint(parameters["publicationID"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPublicationsRepository(db)

	if err = repository.Unlike(publicationID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
