package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"plan2go-backend/config"
	db "plan2go-backend/infra/DB"
	"plan2go-backend/repo"
	"plan2go-backend/rest"
	"plan2go-backend/rest/handlers/activity"
	"plan2go-backend/rest/handlers/guide"
	"plan2go-backend/rest/handlers/plan"
	"plan2go-backend/rest/handlers/user"
	"plan2go-backend/rest/handlers/weather"
	"plan2go-backend/rest/middleware"
	"plan2go-backend/rest/services"

	"google.golang.org/genai"
)

func Serve() {
	dbcn, err := db.ConnectDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Repositories
	userRepo := repo.NewUserRepo(dbcn)
	emailRepo := repo.NewEmailVerificationRepo(dbcn) // <-- new OTP repo
	guideRepo := repo.NewGuideRepo(dbcn)
	activityRepo := repo.NewActivityRepo(dbcn) // <-- new Activity repo

	// Config & Middleware
	cnf := config.GetConfig()
	cnfMiddleWare := middleware.NewConfigMiddleware(cnf)

	// Gemini services
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: os.Getenv("GEMINI_API_KEY"),
	})
	if err != nil {
		log.Fatal(err)
	}
	planServices := services.NewPlanService(client)
	planHandler := plan.NewPlanHandler(planServices)

	guideHandler := guide.NewGuideHandler(guideRepo)
	weatherHandler := weather.NewHandler()
	activityService := services.NewActivityService(activityRepo)
	activityHandler := activity.NewActivityHandler(activityService)

	// User handler with both repos
	userHandler := user.NewHandler(*cnfMiddleWare, userRepo, emailRepo)

	// Start server
	server := rest.NewServer(cnf, userHandler, weatherHandler, planHandler, guideHandler, activityHandler)
	server.Start()
}
