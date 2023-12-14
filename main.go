package main

import (
	"log"
	"net/http"
	"os"
	"pheet-go-line/middleware"

	"github.com/Pheethy/psql/helper"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v8/linebot"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	bot, err := linebot.New(helper.GetENV("CHANNEL_SECRET", ""), helper.GetENV("CHANNEL_TOKEN", ""))
	if err != nil {
		log.Printf("err get bot bot: %v", err)
	}

	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "GoBot จร้าาา")
	})

	r.POST("/callback", func(c *gin.Context) {
		events, err := bot.ParseRequest(c.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				c.Writer.WriteHeader(http.StatusBadRequest)
			} else {
				c.Writer.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if message.Text == "ข้อความ" {
						if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("จะมาข้อความอะไรเนี่ย")).Do(); err != nil {
							log.Println("err:", err)
						}
					} else if message.Text == "สวัสดี" {
						if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ไม่รับบริการอะไรแล้วเหนื่อย")).Do(); err != nil {
							log.Println("err:", err)
						}
					} else if message.Text == "ทำไร" {
						if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ไม่ทำอะไรนะจะนอน")).Do(); err != nil {
							log.Println("err:", err)
						}
					} else {
						if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("พอเถอะง่วงแล้ววว")).Do(); err != nil {
							log.Println("err:", err)
						}
					}
				}

			}
		}
	})

	r.Run(os.Getenv("PORT"))
}
