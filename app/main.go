package main

import (
	"log"
	"sync"

	"sujana-be-web-go/db"
	"sujana-be-web-go/sujana/delivery"
	"sujana-be-web-go/sujana/repository"
	"sujana-be-web-go/sujana/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	Init()
	// initEnv()
	listenPort := ":4001"
	// appName := os.Getenv("APP_NAME")

	usrRepo := repository.NewPostgreUser(db.GormClient.DB)
	accountRepo := repository.NewPostgreAccount(db.GormClient.DB)
	orderRepo := repository.NewPostgreOrder(db.GormClient.DB)
	// opnameRepo := repository.NewPostgreOpname(db.GormClient.DB)

	timeoutContext := fiber.Config{}.ReadTimeout

	userUseCase := usecase.NewUserUseCase(usrRepo, timeoutContext)
	accountUseCase := usecase.NewAccountUseCase(accountRepo, timeoutContext)
	orderUseCase := usecase.NewOrderUseCase(orderRepo, timeoutContext)
	// opnameUseCase := usecase.NewOpnameUseCase(opnameRepo, timeoutContext)

	app := fiber.New(fiber.Config{})
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${green} ${status} ${white} | ${latency} | ${ip} | ${green} ${method} ${white} | ${path} | ${yellow} ${body} ${reset} | ${magenta} ${resBody} ${reset}\n",
		TimeFormat: "02 January 2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))
	app.Use(cors.New())

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {

		//call delivery http here
		delivery.NewUserHandler(app, userUseCase)
		delivery.NewAccountHandler(app, accountUseCase)
		delivery.NewOrderHandler(app, orderUseCase)
		// delivery.NewOpnameHandler(app, opnameUseCase)
		log.Fatal(app.Listen(listenPort))
		wg.Done()
	}()
	wg.Wait()

}

func Init() {
	InitEnv()
	InitDB()
}

func InitEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Println(".env file not found")
	}
}

func InitDB() {
	db.NewGormClient()
}
