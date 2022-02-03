package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/test_kompas/news_app/configs"
	"github.com/test_kompas/news_app/pkg/handler"
	"github.com/test_kompas/news_app/pkg/repository"
	"github.com/test_kompas/news_app/pkg/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db                    = initDB()
	authorRepo            = repository.NewAuthorsRepository(db)
	authenticationRepo    = repository.NewAuthenticationRepository(db)
	articleRepo           = repository.NewArticleRepository(db)
	authorService         = service.NewAuthorsService(authorRepo)
	authenticationService = service.NewAuthenticationService(authenticationRepo, authorRepo)
	articleService        = service.NewArticleService(articleRepo)
	authenticationHandler = handler.NewAuthenticationHandler(authenticationService)
	authorHandler         = handler.NewAuthorsHandler(authenticationService, authorService)
	articleHandler        = handler.NewArticleHandler(authenticationService, articleService)
)

func main() {
	app := fiber.New(fiber.Config{})
	handleArgs()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Authorization, Content-Type, Accept",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	app.Post("/login", authenticationHandler.Login)

	app.Get("/author", authorHandler.Find)

	app.Post("/author", authorHandler.AddAuthor)

	app.Delete("/author", authorHandler.DeleteAuthor)

	app.Get("/article", articleHandler.Find)

	app.Post("/article", articleHandler.AddArticle)

	app.Listen(":8100")
}

func handleArgs() {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			authorIds := seedAuthors()
			seedArticles(authorIds)
		}
	}
}

func initDB() *gorm.DB {

	dsn := configs.GetMySqlDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
