package rest

import (
	"fmt"
	"net/http"
	"os"
	"plan2go-backend/config"
	"plan2go-backend/rest/handlers/activity"
	"plan2go-backend/rest/handlers/guide"
	"plan2go-backend/rest/handlers/plan"
	"plan2go-backend/rest/handlers/user"
	"plan2go-backend/rest/handlers/weather"
	"plan2go-backend/rest/middleware"
	"plan2go-backend/util"
	"strconv"
)

type Server struct {
	cnf             *config.Config
	userHandler     *user.Handler
	weatherHandler  *weather.Handler
	planHandler     *plan.PlanHandler
	guideHandler    *guide.GuideHandler
	activityHandler *activity.ActivityHandler
}

func NewServer(cnf *config.Config,
	userHandler *user.Handler,
	weatherHandler *weather.Handler,
	planHandler *plan.PlanHandler,
	guideHandler *guide.GuideHandler,
	activityHandler *activity.ActivityHandler) *Server {
	return &Server{
		cnf:             cnf,
		userHandler:     userHandler,
		weatherHandler:  weatherHandler,
		planHandler:     planHandler,
		guideHandler:    guideHandler,
		activityHandler: activityHandler,
	}
}

func (server *Server) Start() {
	manager := middleware.NewManager()
	manager.Use(middleware.Logger)
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Backend running 🚀"))
	})

	server.userHandler.RegisterRoutes(mux, manager)
	server.weatherHandler.WeatherRoutes(mux, manager)
	server.planHandler.PlanRoutes(mux, manager)
	server.guideHandler.GuideRoutes(mux, manager)
	server.activityHandler.RegisterActivityRoutes(mux, manager)

	adrr := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("Server is running on port", adrr)
	err := http.ListenAndServe(adrr, util.GlobalRouter(mux))
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
