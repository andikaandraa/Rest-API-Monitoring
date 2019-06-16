package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"

// 	"project/config"
// 	"project/controllers"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/gorilla/mux"
// 	"github.com/prometheus/client_golang/prometheus"
// 	"github.com/prometheus/client_golang/prometheus/promauto"
// 	"github.com/prometheus/client_golang/prometheus/promhttp"
// )

// func recordMetrics() {
// 	go func() {
// 		for {
// 			opsProcessed.Inc()
// 			time.Sleep(2 * time.Second)
// 		}
// 	}()
// }

// var (
// 	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
// 		Name: "myapp_processed_ops_total",
// 		Help: "The total number of processed events",
// 	})
// )

// func handleRequest(port string) {
// 	db := config.DBInit()
// 	objDB := &controllers.ObjDB{DB: db}

// 	router := mux.NewRouter()
// 	router.HandleFunc("/", controllers.HomepageHandler)
// 	router.HandleFunc("/user", objDB.HandleUserRequest)
// 	router.HandleFunc("/user/", objDB.HandleGetUserRequest)
// 	router.HandleFunc("/users", objDB.HandleUsersRequest)
// 	router.Handle("/metrics", promhttp.Handler())

// 	fmt.Println("Serve at port ", port)
// 	log.Fatal(http.ListenAndServe(port, router))
// }

// func main() {
// 	port := ":9001"
// 	recordMetrics()
// 	handleRequest(port)
// 	fmt.Println("Hello, playground")
// }
