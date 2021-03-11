package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/menza01/model"
	"github.com/menza01/views"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func todo() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if r.Method == http.MethodPost {

			data := views.TodoRequest{
				Name:   "",
				Todo:   "",
				Status: -1,
			}

			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				errData := views.ErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				}
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(errData)
				return
			}

			if data.Name == "" || data.Todo == "" || data.Status == -1 {
				errData := views.ErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: "Field invalidate",
				}
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(errData)
				return
			}

			fmt.Println(data)
			if err := model.CreateTodo(data.Name, data.Todo, data.Status); err != nil {
				errData := views.ErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				}
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(errData)
				return
			}

			w.WriteHeader(http.StatusCreated)
			dataResponse := views.Response{
				Code: http.StatusOK,
				Body: data,
			}
			json.NewEncoder(w).Encode(dataResponse)
		} else if r.Method == http.MethodGet {

			data, err := model.ReadAll()
			if err != nil {
				errData := views.ErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				}
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(errData)
				return
			}

			w.WriteHeader(http.StatusOK)
			dataResponse := views.Response{
				Code: http.StatusOK,
				Body: data,
			}
			json.NewEncoder(w).Encode(dataResponse)
		} else if r.Method == http.MethodPut {
			data := views.TodoUpdate{
				Id:     -1,
				Status: -1,
				Name:   "",
				Todo:   "",
			}

			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				errData := views.ErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				}
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(errData)
				return
			}

			if data.Id == -1 || data.Status == -1 || data.Name == "" || data.Todo == "" {
				errData := views.ErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: "Field invalidate",
				}
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(errData)
				return
			}

			fmt.Println(data)
			if err := model.UpdateTodo(data.Status, data.Id, data.Name, data.Todo); err != nil {
				errData := views.ErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				}
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(errData)
				return
			}
			w.WriteHeader(http.StatusOK)

			dataResponse := views.Response{
				Code: http.StatusOK,
				Body: data,
			}
			json.NewEncoder(w).Encode(dataResponse)

		} else if r.Method == http.MethodDelete {
			// todo
			data := views.TodoDelete{
				Id: -1,
			}
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				errData := views.ErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				}
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(errData)
				return
			}

			if data.Id == -1 {
				errData := views.ErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: "Field invalidate",
				}
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(errData)
				return
			}
			fmt.Println(data)
			if err := model.DeleteTodo(data.Id); err != nil {
				errData := views.ErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				}
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(errData)
				return
			}
			w.WriteHeader(http.StatusOK)

			dataResponse := views.Response{
				Code: http.StatusOK,
				Body: data,
			}
			json.NewEncoder(w).Encode(dataResponse)
		}
	}
}
