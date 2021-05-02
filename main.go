package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/websocket/v2"
	"log"
	"os"
)

func hello(ctx *fiber.Ctx) error{
	res := fmt.Sprintf("server-%s" , os.Args[2])
	return ctx.SendString(res)
}

func main ()  {
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/" , hello)
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

	port := fmt.Sprintf(":%s" , os.Args[1])

	log.Fatal(app.Listen(port))
}


