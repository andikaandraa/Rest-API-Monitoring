package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"project/config"
	"project/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	circuit "github.com/rubyist/circuitbreaker"
	"github.com/shirou/gopsutil/mem"
)

var (
	cb              *circuit.Breaker
	total           int
	successRequests float64
	failureRequests float64
)

// ObjDB represent object database
type ObjDB controllers.ObjDB

func HandleError(w http.ResponseWriter, r *http.Request) {
	fmt.Println(total)
}

// HandleRoute handle route for breaker
func (objDB *ObjDB) HandleRoute(w http.ResponseWriter, r *http.Request) {
	cb.Call(func() error {
		var err error
		switch r.URL.RequestURI() {
		case "/user":
			err = controllers.HandleUserRequest(w, r, objDB.DB)
		case "/users":
			err = controllers.HandleUsersRequest(w, r, objDB.DB)
		default:
			err = controllers.HandleGetUserRequest(w, r, objDB.DB)
		}
		if err != nil {
			total = total + 1
			failureRequests = failureRequests + 1
		} else {
			successRequests = successRequests + 1
		}
		return err
	}, 0)
}

func handleRequest(port string) {
	cb = circuit.NewConsecutiveBreaker(10)
	total = 0
	events := cb.Subscribe()
	go func() {
		for {
			e := <-events
			fmt.Println("-->", e)
			// Monitor breaker events like BreakerTripped, BreakerReset, BreakerFail, BreakerReady
		}
	}()

	logger := log.New(os.Stdout, "[Memory]", log.Lshortfile|log.Ldate|log.Ltime)
	db := config.DBInit()
	objDB := &ObjDB{DB: db}

	router := mux.NewRouter()

	router.HandleFunc("/", controllers.HomepageHandler)
	router.HandleFunc("/error", HandleError)
	router.HandleFunc("/user", objDB.HandleRoute)
	router.HandleFunc("/user/{id}", objDB.HandleRoute)
	router.HandleFunc("/users", objDB.HandleRoute)
	router.Handle("/metrics", promhttp.Handler())

	// log.Fatal(http.ListenAndServe(port, router))
	memoryPercent := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "memory_percent",
		Help: "memory use percent",
	},
		[]string{"percent"},
	)

	requestsRate := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "requests_rate",
		Help: "requests rate",
	},
		[]string{"requests"},
	)
	prometheus.MustRegister(requestsRate)

	go func() {
		http.ListenAndServe(port, router)
	}()

	for {
		v, err := mem.VirtualMemory()
		if err != nil {
			fmt.Println("get memory use percent error:%s", err)
			logger.Println("get memory use percent:")

		}
		usedPercent := v.UsedPercent
		// logger.Println("get memory use percent:", usedPercent)
		memoryPercent.WithLabelValues("usedMemory").Set(usedPercent)
		requestsRate.WithLabelValues("successRequests").Set(successRequests)
		requestsRate.WithLabelValues("failureRequests").Set(failureRequests)

		time.Sleep(time.Second * 2)
	}

}

func main() {
	port := ":9090"
	handleRequest(port)
}
