package lambo_bot

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func PingHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	resp := fmt.Sprintf("System is up. %s", time.Now().Format("2006-01-02T15:04:05-0700"))
	w.Write([]byte(resp))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("not using .env file")
	}

	router := httprouter.New()
	router.GET("/ping", PingHandler)
	router.NotFound = http.FileServer(http.Dir("static"))

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	log.Printf("listening on port: %s", port)
	serviceAddress := fmt.Sprintf(":%s", port)
	srv := &http.Server{
		Addr:              serviceAddress,
		Handler:           router,
	}
	log.Fatal(srv.ListenAndServe())
}
