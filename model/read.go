package model

import "github.com/menza01/views"

func ReadAll() ([]views.TodoReponse, error) {
	rows, err := con.Query("SELECT * FROM TODO")
	if err != nil {
		return nil, err
	}

	todos := []views.TodoReponse{}
	for rows.Next() {
		data := views.TodoReponse{}
		rows.Scan(&data.Id, &data.Name, &data.Todo, &data.Status)
		todos = append(todos, data)
	}
	return todos, nil
}
