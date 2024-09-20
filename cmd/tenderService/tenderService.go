package main

import (
	"log"
	"net/http"
	"os"
	"tenderService/internal/db"
	"tenderService/internal/handlers"
)

func main() {
	err:= db.Init()
	if err != nil {
		log.Fatalf("Database initialisation error: %v", err)
	}
	
	serverAddr, exist := os.LookupEnv("SERVER_ADDRESS")
	if !exist {
		log.Fatal("SERVER_ADDRESS env. variable doesn't exist!")
	}

	http.HandleFunc("GET /api/ping", handlers.HandlePing)
	http.HandleFunc("GET /api/tenders", handlers.HandleTendersGet)
	http.HandleFunc("POST /api/tenders/new", handlers.HandleTenderNew)
	http.HandleFunc("GET /api/tenders/my", handlers.HandleTenderMy)
	http.HandleFunc("GET /api/tenders/{tenderId}/status", handlers.HandleTenderGetStatus)
	http.HandleFunc("PUT /api/tenders/{tenderId}/status", handlers.HandleTenderEditStatus)
	http.HandleFunc("PATCH /api/tenders/{tenderId}/edit", handlers.HandleTenderEdit)
	http.HandleFunc("POST /api/bids/new", handlers.HandleCreateBid)                            // POST: Создание нового предложения
	http.HandleFunc("GET /api/bids/my", handlers.HandleGetMyBids)                              // GET: Получение списка ваших предложений
	http.HandleFunc("GET /api/bids/{tenderId}/list", handlers.HandleGetBidsByTender)           // GET: Получение списка предложений для тендера
	http.HandleFunc("GET /api/bids/{bidId}/status", handlers.HandleGetBidStatus)               // GET: Получение текущего статуса предложения
	http.HandleFunc("PUT /api/bids/{bidId}/status", handlers.HandleEditBidStatus)              // PUT: Изменение статуса предложения
	http.HandleFunc("PATCH /api/bids/{bidId}/edit", handlers.HandleEditBid)                    // PATCH: Редактирование параметров предложения
	http.HandleFunc("PUT /api/bids/{bidId}/submit_decision", handlers.HandleSubmitBidDecision) // PUT: Отправка решения по предложению

	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
