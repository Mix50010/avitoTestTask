package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"tenderService/internal/db"
	"tenderService/internal/models"
)

func HandleTendersGet(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing TendersGet request...")

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

	// Обработка параметра "service_type"
	var serviceTypes []models.TenderServiceType
	serviceTypesStr := query["service_type"] // Возвращает массив значений
	if len(serviceTypesStr) > 0 {
		log.Println("Query has service_type parameter!")

		// Валидация service_type
		for _, serviceType := range serviceTypesStr {
			serviceTypeCheck := models.TenderServiceType(serviceType)
			if !serviceTypeCheck.IsValid() {
				WriteError(w, "Service type isn't recognisable", http.StatusBadRequest)
				return
			}
		}

		for _, serviceType := range serviceTypesStr {
			serviceTypes = append(serviceTypes, models.TenderServiceType(serviceType))
		}
	}

	log.Printf("Limit: %d, Offset: %d, ServiceTypes: %v", limitParam, offsetParam, serviceTypes)


	tenders, err := db.GetTenders(limitParam, offsetParam, serviceTypesStr)
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(tenders)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func HandleTenderNew(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing TenderNew request...")

	req := models.CreateTenderRequest{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&req); err != nil {
		WriteError(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, err := db.GetIDByUsername(req.CreatorUsername)
	if err != nil {
		WriteError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	isResponsible, err := db.IsResponsible(userId, req.OrganizationId)
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !isResponsible {
		WriteError(w, "User hasn't rights", http.StatusForbidden)
		return
	}

	tenderID, createdAt, err := db.NewTender(req.Name, req.Description, string(req.ServiceType), "Created", req.OrganizationId, "1")
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.AddTenderToUser(userId, tenderID)
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusOK)
	data, err := json.Marshal(models.Tender{
		Id:             tenderID,
		Name:           req.Name,
		Description:    req.Description,
		ServiceType:    req.ServiceType,
		Status:         "Created",
		OrganizationId: req.OrganizationId,
		Version:        1,
		CreatedAt:      createdAt,
	})
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func HandleTenderMy(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing TenderMy request...")

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

	tenders, err := db.GetUserTenders(userID, limitParam, offsetParam)
	if err != nil {
		WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(tenders)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}


func HandleTenderGetStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing TenderGetStatus request...")

	tenderID := r.PathValue("tenderId")
	if len(tenderID) == 0 {
		WriteError(w, "Tender ID isn't recognisable", http.StatusBadRequest)
		return
	}

	query := r.URL.Query()

	var username string
	if query.Has("username") {
		log.Println("Query has username parameter!")
		username = query.Get("username")
	} else {
		WriteError(w, "Username isn't recognisable", http.StatusUnauthorized)
		return
	}

	userID, err := db.GetIDByUsername(username)
	if err != nil {
		WriteError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	tender, err := db.GetTenderByID(tenderID)
	if err != nil {
		WriteError(w, err.Error(), http.StatusNotFound)
		return
	}
	orgID := tender.OrganizationId

	isResponsible, err := db.IsResponsible(userID, orgID)
	if err != nil {
		WriteError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if !isResponsible {
		WriteError(w, "User hasn't rights", http.StatusForbidden)
		return
	}

	status, err := db.GetTenderStatus(tenderID)
	if err != nil {
		WriteError(w, err.Error(), http.StatusNotFound)
		return
	}
	data, err := json.Marshal(status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func HandleTenderEditStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing TenderEditStatus request...")

	tenderID := r.PathValue("tenderId")
	if len(tenderID) == 0 {
		WriteError(w, "Tender ID isn't recognisable", http.StatusBadRequest)
		return
	}

	query := r.URL.Query()

	var username string
	if query.Has("username") {
		log.Println("Query has username parameter!")
		username = query.Get("username")
	} else {
		WriteError(w, "Username isn't recognisable", http.StatusBadRequest)
		return
	}

	var status string
	if query.Has("status") {
		log.Println("Query has status parameter!")
		status = query.Get("status")
	} else {
		WriteError(w, "Status isn't recognisable", http.StatusBadRequest)
		return
	}

	userID, err := db.GetIDByUsername(username)
	if err != nil {
		WriteError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	tender, err := db.GetTenderByID(tenderID)
	if err != nil {
		WriteError(w, err.Error(), http.StatusNotFound)
		return
	}
	orgID := tender.OrganizationId

	isResponsible, err := db.IsResponsible(userID, orgID)
	if err != nil {
		WriteError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if !isResponsible {
		WriteError(w, "User hasn't rights", http.StatusForbidden)
		return
	}

	err = db.EditTenderStatus(tenderID, status)
	if err != nil {
		WriteError(w, err.Error(), http.StatusNotFound)
		return
	}
	data, err := json.Marshal(status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func HandleTenderEdit(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing TenderEdit request...")
}
