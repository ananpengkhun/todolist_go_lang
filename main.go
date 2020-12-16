package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/menza01/controller"
	"github.com/menza01/model"
)

type customer struct {
	Firname  string `json:"first"`
	Lastname string `json:"last"`
	Age      int    `json:"a"`
	Sex      string `json:"s"`
}

func main() {

	customerLust := []customer{}

	customer1 := customer{
		Firname:  "men",
		Lastname: "peng",
		Age:      22,
		Sex:      "M",
	}

	customer2 := customer{
		Firname:  "men2",
		Lastname: "peng2",
		Age:      22,
		Sex:      "M",
	}

	customerLust = append(customerLust, customer1)
	customerLust = append(customerLust, customer2)

	data, err := json.Marshal(&customerLust)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	http.ListenAndServe(":3000", mux)

}
