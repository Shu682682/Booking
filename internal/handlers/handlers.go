package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Shu682682/Booking.git/internal/config"
	"github.com/Shu682682/Booking.git/internal/forms"
	"github.com/Shu682682/Booking.git/internal/models"
	"github.com/Shu682682/Booking.git/internal/render"
)

//Repo the repository used by the handlers
var Repo *Repository
//Repository is the repository type
type Repository struct{
	App *config.AppConfig

}
//New Repor creats a new repository
func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App:a,
	}
}

//NewHandlers sets the repoitory for the handlers
func NewHandlers(r *Repository){
	Repo=r
}

//Home is the home page handler
func (m *Repository)Home(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r, "home.html", &models.TemplateData{})
	
}
//About is the about page handler
func(m *Repository) About(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r,"about.html", &models.TemplateData{} )
}

func(m*Repository)Index(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r,"index.html", &models.TemplateData{})

}
func(m*Repository)General(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r,"generals_quarter.html", &models.TemplateData{})

}
func(m*Repository)Major(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r,"majors_suite.html", &models.TemplateData{})

}

func(m*Repository)Contact(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r,"contact.html", &models.TemplateData{})

}
//Get Availability renders the search availability page
func(m*Repository)Book(w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, r,"book.html", &models.TemplateData{})

}
func(m*Repository)PostBook(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Invalid request method",
		})
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing form:", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Unable to parse form",
		})
		return
	}

	//gain users' data
	start := r.FormValue("start_date")
	end := r.FormValue("end_date")
	people := r.FormValue("people_amount")
	name := r.FormValue("full_name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	room := r.FormValue("room_choice")

	
	if start == "" || end == "" || people == "" || name == "" || email == "" || phone == "" || room == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": "Missing required fields",
		})
		return
	}

	// response json
	response := map[string]interface{}{
		"success":       true,
		"message":       "Booking confirmed!",
		"start_date":    start,
		"end_date":      end,
		"people_amount": people,
		"full_name":     name,
		"email":         email,
		"phone":         phone,
		"room_choice":   room,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}


type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json") 

    err := r.ParseForm()
    if err != nil {
        log.Println("Error parsing form:", err)
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(jsonResponse{
            OK:      false,
            Message: "Unable to parse form data",
        })
        return
    }

    start := r.Form.Get("start")  // 👈 改成 "start"
    end := r.Form.Get("end")      // 👈 改成 "end"

    if start == "" || end == "" {
        log.Println("Missing required fields")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(jsonResponse{
            OK:      false,
            Message: "Missing required fields",
        })
        return
    }

    log.Printf("Checking availability from %s to %s", start, end)

    w.WriteHeader(http.StatusOK) 
    json.NewEncoder(w).Encode(jsonResponse{
        OK:      true,
        Message: "Available!",
    })
}

// func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json") 

//     err := r.ParseForm()
//     if err != nil {
//         w.WriteHeader(http.StatusBadRequest)
//         json.NewEncoder(w).Encode(jsonResponse{
//             OK:      false,
//             Message: "Unable to parse form data",
//         })
//         return
//     }

//     start := r.Form.Get("start_date")
//     end := r.Form.Get("end_date")

//     if start == "" || end == "" {
//         w.WriteHeader(http.StatusBadRequest)
//         json.NewEncoder(w).Encode(jsonResponse{
//             OK:      false,
//             Message: "Missing required fields",
//         })
//         return
//     }

//     log.Printf("Checking availability from %s to %s", start, end)

//     w.WriteHeader(http.StatusOK) 
//     json.NewEncoder(w).Encode(jsonResponse{
//         OK:      true,
//         Message: "Available!",
//     })
// }

//Reservation renders the make a reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
    form := forms.New(nil) // initialize Form
    data := models.TemplateData{
        Form: form, // Pass Form for tempalte
    }

    render.RenderTemplate(w, r, "book.html", &data)
}

//PostReservation handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing form:", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Unable to parse form",
		})
		return
	}

	// 取得 `people_amount`
	peopleStr := r.Form.Get("people_amount")
	people, err := strconv.Atoi(peopleStr)
	if err != nil {
		log.Println("Invalid people amount:", peopleStr, err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Invalid number of people",
		})
		return
	}

	reservation := models.Reservation{
		StartDate:  r.Form.Get("start_date"),
		EndDate:    r.Form.Get("end_date"),
		People:     people,
		FullName:   r.Form.Get("full_name"),
		Email:      r.Form.Get("email"),
		Phone:      r.Form.Get("phone"),
		RoomChoice: r.Form.Get("room_choice"),
	}

	missingFields := []string{}
	if reservation.StartDate == "" { missingFields = append(missingFields, "start_date") }
	if reservation.EndDate == "" { missingFields = append(missingFields, "end_date") }
	if reservation.FullName == "" { missingFields = append(missingFields, "full_name") }
	if reservation.Email == "" { missingFields = append(missingFields, "email") }
	if reservation.Phone == "" { missingFields = append(missingFields, "phone") }
	if reservation.RoomChoice == "" { missingFields = append(missingFields, "room_choice") }

	if len(missingFields) > 0 {
		log.Println("Missing fields:", missingFields)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error":   "Missing required fields",
			"fields":  missingFields,
		})
		return
	}

	log.Println("Reservation received:", reservation)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":     true,
		"start_date":  reservation.StartDate,
		"end_date":    reservation.EndDate,
		"people":      reservation.People,
		"full_name":   reservation.FullName,
		"email":       reservation.Email,
		"phone":       reservation.Phone,
		"room_choice": reservation.RoomChoice,
		"message":     "Booked Successfully!",
	})
}
