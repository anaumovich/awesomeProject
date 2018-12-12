package OldFolder

import (
	hello "awesomeProject/OldFolder/proto"
	"context"
	_ "fmt"
	"log"
	"net/http"

	pb "./proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

var grpcClient hello.NotificatorClient

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/home/", http.StripPrefix("/home/", fs))

	http.HandleFunc("/", HomeRouterHandler)
	error := http.ListenAndServe(":9000", nil)
	if error != nil {
		log.Fatal("ListenAndServe: ", error)
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	grpcClient = pb.NewNotificatorClient(conn)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

}

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		_ = r.ParseForm()
		_ = r.Form["login"]
		_ = r.Form["password"]
		http.Redirect(w, r, "http://localhost:9000/home", http.StatusSeeOther)
		_, _ = grpcClient.SendMessage(context.Background(), &hello.MessageReq{Type: "Warning", Text: "Service restarted"})
	}
}
