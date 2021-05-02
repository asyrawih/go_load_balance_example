package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/websocket/v2"
	"log"
	"os"
)
// hello ctx
func hello(ctx *fiber.Ctx) error{
	res := fmt.Sprintf("server-%s" , os.Args[2])
	return ctx.SendString(res)
}

func main ()  {
	// Calling Fiber func
	app := fiber.New()
	// Generate Logger
	app.Use(logger.New())
	// Create Root Endpoint
	app.Get("/" , hello)
	// Upgrade Http Allow Websocket
	app.Get("/ws" , websocket.New(func(conn *websocket.Conn) {
		fmt.Println(conn.Locals("Host"))
		for  {
			mt , msg , err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)
			err = conn.WriteMessage(mt, msg)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}))
	// Accept Argument on Command Line
	port := fmt.Sprintf(":%s" , os.Args[1])
	// running server
	log.Fatal(app.Listen(port))
}


