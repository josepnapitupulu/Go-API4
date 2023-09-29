package routes

import (
	"github.com/gorilla/mux"
	"github.com/josepnapitupulu/Project_pasti/controllers"
)

var RegisterJemaatsRoutes = func(router *mux.Router) {
	// router.HandleFunc("/", controllers.Index).Methods("GET")
	// router.HandleFunc("/jemaatBaru", controllers.Create).Methods("POST")
	router.HandleFunc("/jemaat/", controllers.CreateJemaat).Methods("POST")
	router.HandleFunc("/jemaat/", controllers.GetJemaat).Methods("GET")
	router.HandleFunc("/jemaat/{jemaatId}", controllers.GetJemaatById).Methods("GET")
	router.HandleFunc("/jemaat/{jemaatId}", controllers.UpdateJemaat).Methods("PUT")
	router.HandleFunc("/jemaat/{jemaatId}", controllers.DeleteJemaat).Methods("DELETE")
}
