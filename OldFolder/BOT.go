package OldFolder

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "./proto"
	"google.golang.org/grpc"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

const (
	BotToken = "763230126:AAFOyQADuNsVw2GeHv3uYL-xPNjEJygZsv8"
	/*WebHookUrl = "https://b2cb2901.ngrok.io"*/
	port = ":50051"
)

type server struct {
	Bot *tgbotapi.BotAPI
}

/*var button1 = tgbotapi.NewInlineKeyboardButtonData("Start", "/start@AntohBot")
var button2 = tgbotapi.NewInlineKeyboardButtonData("About!", "/About@AntohBot")
var button3 = tgbotapi.NewInlineKeyboardButtonData("Hello!", "/Hello@AntohBot")
var row = tgbotapi.NewInlineKeyboardRow(button1, button2, button3)
var keyboard = tgbotapi.NewInlineKeyboardMarkup(row)
*/
func main() {
	bot, err := tgbotapi.NewBotAPI(BotToken)

	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	serverGRPC := &server{bot}
	pb.RegisterNotificatorServer(s, serverGRPC)

	go http.ListenAndServe(":8080", nil)
	fmt.Println("start listen :8080")

	/*updates := bot.ListenForWebhook("/")

	for update := range updates {

		if update.Message != nil {
			home := update.Message.Text

			if home == "/start@AntohBot" || home == "/start" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "choose  you command")
				msg.ReplyMarkup = keyboard
				_, _ = bot.Send(msg)
			}
		}

		if update.CallbackQuery != nil {
			textData := update.CallbackQuery.Data

			if textData == "/start@AntohBot" {
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "choose  you command")
				msg.ReplyMarkup = keyboard
				_, _ = bot.Send(msg)
			}
			if textData == "/About@AntohBot" {
				_, _ = bot.Send(tgbotapi.NewMessage(
					update.CallbackQuery.Message.Chat.ID,
					"Hi, i'm a Anton bot, let's start.",
				))
			} else if textData == "/Hello@AntohBot" {
				str, err := s.takeMessage()
				if err != nil {
					_, _ = bot.Send(tgbotapi.NewMessage(
						update.CallbackQuery.Message.Chat.ID,
						"sorry, error happened",
					))
				} else {
					_, _ = bot.Send(tgbotapi.NewMessage(
						update.CallbackQuery.Message.Chat.ID,
						str,
					))
				}
			}
		}
	}*/
}

func (s *server) SendMessage(ctx context.Context, in *pb.MessageReq) (*pb.Empty, error) {
	s.Bot.Send()
	return &pb.Empty{}, nil
}
