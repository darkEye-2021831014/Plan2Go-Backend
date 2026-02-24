// cmd/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"plan2go-backend/config"
	db "plan2go-backend/infra/DB"
	"plan2go-backend/repo"
	"plan2go-backend/rest/handlers/activity"
	"plan2go-backend/rest/handlers/guide"
	"plan2go-backend/rest/handlers/plan"
	"plan2go-backend/rest/handlers/user"
	"plan2go-backend/rest/handlers/weather"
	"plan2go-backend/rest/middleware"
	"plan2go-backend/rest/services"
	"plan2go-backend/util"

	"github.com/rs/cors"
	"google.golang.org/genai"
)

func main() {
	// ---- Connect Database ----
	dbcn, err := db.ConnectDB()
	if err != nil {
		fmt.Println("DB connection error:", err)
		os.Exit(1)
	}

	// ---- Repositories ----
	userRepo := repo.NewUserRepo(dbcn)
	emailRepo := repo.NewEmailVerificationRepo(dbcn)
	guideRepo := repo.NewGuideRepo(dbcn)
	activityRepo := repo.NewActivityRepo(dbcn)

	// ---- Config & Middleware ----
	cnf := config.GetConfig()
	cnfMiddleware := middleware.NewConfigMiddleware(cnf)

	// ---- Gemini Services ----
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: os.Getenv("GEMINI_API_KEY"),
	})
	if err != nil {
		log.Fatal("Gemini client error:", err)
	}

	planServices := services.NewPlanService(client)
	planHandler := plan.NewPlanHandler(planServices)
	guideHandler := guide.NewGuideHandler(guideRepo)
	weatherHandler := weather.NewHandler()
	activityService := services.NewActivityService(activityRepo)
	activityHandler := activity.NewActivityHandler(activityService)
	userHandler := user.NewHandler(*cnfMiddleware, userRepo, emailRepo)

	// ---- Setup Router ----
	manager := middleware.NewManager()
	manager.Use(middleware.Logger)
	mux := http.NewServeMux()
	userHandler.RegisterRoutes(mux, manager)
	weatherHandler.WeatherRoutes(mux, manager)
	planHandler.PlanRoutes(mux, manager)
	guideHandler.GuideRoutes(mux, manager)
	activityHandler.RegisterActivityRoutes(mux, manager)

	// ---- Enable CORS for frontend ----
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // replace "*" with your Vercel frontend URL in prod
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})
	handler := c.Handler(util.GlobalRouter(mux))

	// ---- Use Railway PORT or fallback ----
	port := os.Getenv("PORT")
	if port == "" {
		port = fmt.Sprintf("%d", cnf.HttpPort) // local fallback
	}
	fmt.Println("Server is running on port", port)

	// ---- Start HTTP Server ----
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
