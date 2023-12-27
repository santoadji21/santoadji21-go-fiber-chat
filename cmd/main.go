package main

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/santoadji21/santoadji21-go-fiber-chat/internal/db"
	"github.com/santoadji21/santoadji21-go-fiber-chat/internal/redis"
	"github.com/santoadji21/santoadji21-go-fiber-chat/pkg/config"
	"github.com/santoadji21/santoadji21-go-fiber-chat/pkg/handler"
)

func main() {

	//  App Static
	// Create a new engine
	engine := html.New("/app/web/static", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", fiber.Map{})
	})

	// WebSocket route for chat
	app.Get("/ws", websocket.New(handler.WebSocketChatHandler))

	// Load your configuration
	cfg := config.DbCfg()

	// Initialize the database
	db.ConnectDB(cfg)

	// Initialize Redis
	redis.ConnectRedis(cfg)

	// Check if the database connection was successful
	if db.GetDB() != nil {
		log.Println("Successfully connected to the database")
	} else {
		log.Fatalln("Failed to connect to the database")
	}

	// Check if the Redis connection was successful
	if redis.GetRedis() != nil {
		log.Println("Successfully connected to Redis")
	} else {
		log.Fatalln("Failed to connect to Redis")
	}
	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
