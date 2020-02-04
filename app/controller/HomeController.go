package controller

import (
	"net/http"

	"github.com/chazeprasad/go-api-seed/app/util"
)

func Home(w http.ResponseWriter, r *http.Request) {
	util.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
