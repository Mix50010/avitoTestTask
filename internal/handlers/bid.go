package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"tenderService/internal/db"
	"tenderService/internal/models"
)

// Обработчик для создания нового предложения
func HandleCreateBid(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing CreateBid request...")

	req := models.CreateBidRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&req); err != nil {
		WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}
	switch req.AuthorType {

	case models.ORGANIZATION:
		_, err := db.GetTenderByID(req.TenderId)
		if err != nil {
			WriteError(w, err.Error(), http.StatusNotFound)
			return
		}
	case models.USER:
		tender, err := db.GetTenderByID(req.TenderId)
		if err != nil {
			WriteError(w, err.Error(), http.StatusNotFound)
			return
		}

		isResponsible, err := db.IsResponsible(req.AuthorId, tender.OrganizationId)
		if err != nil {
			WriteError(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if !isResponsible {
			WriteError(w, "User hasn't rights", http.StatusForbidden)
			return
		}
	default:
		WriteError(w, "unexpected models.BidAuthorType", http.StatusBadRequest)
		return
	}

	bidID, createdAt, err := db.CreateBid(req)
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(models.Bid{
		Id:             bidID,
		Name:           req.Name,
		Description:    req.Description,
		Status:         "Created",
		TenderId:       req.TenderId,
		AuthorType:     req.AuthorType,
		AuthorId:       req.AuthorId,
		Version:        1,
		CreatedAt:      createdAt,
	})
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)

}

// Обработчик для получения списка ваших предложений
func HandleGetMyBids(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing GetMyBids request...")

	query := r.URL.Query()

	// Обработка параметра "limit"
	var limitParam int
	if query.Has("limit") {
		log.Println("Query has limit parameter!")
		var err error
		limit, err := strconv.Atoi(query.Get("limit"))
		if err != nil {
			WriteError(w, "Limit isn't recognisable", http.StatusBadRequest)
			return
		}
		limitParam = limit
	} else {
		limitParam = 5
	}

	// Обработка параметра "offset"
	var offsetParam int
	if query.Has("offset") {
		log.Println("Query has offset parameter!")
		var err error
		offset, err := strconv.Atoi(query.Get("offset"))
		if err != nil {
			WriteError(w, "Offset isn't recognisable", http.StatusBadRequest)
			return
		}
		offsetParam = offset
	} else {
		offsetParam = 0
	}
	
	var username string
	if query.Has("username") {
		log.Println("Query has username parameter!")
		username = query.Get("username")
	} else {
		WriteError(w, "Username isn't recognisable", http.StatusUnauthorized)
		return
	}

	log.Printf("Limit: %d, Offset: %d, Username: %v", limitParam, offsetParam, username)

	userID, err := db.GetIDByUsername(username)
	if err != nil {
		WriteError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	bids, err := db.GetBidsByAuthorID(userID, limitParam, offsetParam)
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(bids)
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

// Обработчик для получения списка предложений для конкретного тендера
func HandleGetBidsByTender(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing GetBidsByTender request...")

	tenderID := r.PathValue("tenderId")
	if len(tenderID) == 0 {
		WriteError(w, "Tender ID isn't recognisable", http.StatusBadRequest)
		return
	}

	query := r.URL.Query()

	// Обработка параметра "limit"
	var limitParam int
	if query.Has("limit") {
		log.Println("Query has limit parameter!")
		var err error
		limit, err := strconv.Atoi(query.Get("limit"))
		if err != nil {
			WriteError(w, "Limit isn't recognisable", http.StatusBadRequest)
			return
		}
		limitParam = limit
	} else {
		limitParam = 5
	}

	// Обработка параметра "offset"
	var offsetParam int
	if query.Has("offset") {
		log.Println("Query has offset parameter!")
		var err error
		offset, err := strconv.Atoi(query.Get("offset"))
		if err != nil {
			WriteError(w, "Offset isn't recognisable", http.StatusBadRequest)
			return
		}
		offsetParam = offset
	} else {
		offsetParam = 0
	}

	var username string
	if query.Has("username") {
		log.Println("Query has username parameter!")
		username = query.Get("username")
	} else {
		WriteError(w, "Username isn't recognisable", http.StatusUnauthorized)
		return
	}

	log.Printf("Limit: %d, Offset: %d, Username: %v, TenderID: %v", limitParam, offsetParam, username, tenderID)

	tenders, err := db.GetBidsByTenderID(tenderID, limitParam, offsetParam)
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(tenders)
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

// Обработчик для получения текущего статуса предложения
func HandleGetBidStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing GetBidStatus request...")

	query := r.URL.Query()

	bidID := r.PathValue("bidId")
	if len(bidID) == 0 {
		WriteError(w, "Bid ID isn't recognisable", http.StatusBadRequest)
		return
	}

	var username string
	if query.Has("username") {
		log.Println("Query has username parameter!")
		username = query.Get("username")
	} else {
		WriteError(w, "Username isn't recognisable", http.StatusUnauthorized)
		return
	}

	log.Printf("BidID: %v, Username: %v", bidID, username)
	//TODO
	_, err := db.GetIDByUsername(username)
	if err != nil {
		WriteError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	status, err := db.GetBidStatusByID(bidID)
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(status)
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

// Обработчик для изменения статуса предложения
func HandleEditBidStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing EditBidStatus request...")

	query := r.URL.Query()

	bidID := r.PathValue("bidId")
	if len(bidID) == 0 {
		WriteError(w, "Bid ID isn't recognisable", http.StatusBadRequest)
		return
	}

	var username string
	if query.Has("username") {
		log.Println("Query has username parameter!")
		username = query.Get("username")
	} else {
		WriteError(w, "Username isn't recognisable", http.StatusUnauthorized)
		return
	}

	var status string
	if query.Has("status") {
		log.Println("Query has status parameter!")
		status = query.Get("status")
	} else {
		WriteError(w, "Status isn't recognisable", http.StatusUnauthorized)
		return
	}

	log.Printf("BidID: %v, Username: %v, Status: %v", bidID, username, status)
	//TODO
	_, err := db.GetIDByUsername(username)
	if err != nil {
		WriteError(w, err.Error(), http.StatusUnauthorized)
		return
	}


	err = db.SetBidStatusByID(bidID, status)
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bid, err := db.GetBidByID(bidID)
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(bid)
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

// Обработчик для редактирования параметров предложения
func HandleEditBid(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing EditBid request...")
}

// Обработчик для отправки решения по предложению
func HandleSubmitBidDecision(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing SubmitBidDecision request...")
}
