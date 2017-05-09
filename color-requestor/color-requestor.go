package main

// Import packages
import (
	"flag"
	"log"
	"net/http"

	"strings"

	personProto "charlotte/common/person/v1"

	"github.com/willcj33/hello-kube/color-requestor/grpcHandler"
	socket "github.com/willcj33/hello-kube/color-requestor/socket"
)

var forever = make(chan struct{})

func main() {
	flag.Parse()
	grpcClient, conn := grpcHandler.NewClient()
	defer conn.Close()
	hub := socket.NewHub(grpcClient)
	go hub.Run()
	http.HandleFunc("/", renderIndexPage)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		socket.SocketHandler(hub, w, r)
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	newthing = personProto.Person{}

	/*tickerLoad := time.NewTicker(10 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-tickerLoad.C:
				color, err := client.GetColor(context.Background(), &pb.Empty{})
				if err != nil {
					log.Println(err)
					break
				}
				log.Printf("Color: %v\n", color.Hex)
			case <-quit:
				tickerLoad.Stop()
			}
		}
	}()*/
	<-forever
}

func renderIndexPage(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	if strings.TrimSpace(r.URL.Path) != "/" {
		http.Error(w, "Not found", 404)
		log.Println("in 404")
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	http.ServeFile(w, r, "index.html")
}
