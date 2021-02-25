package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	bnkgopackage "github.com/chaiyarin/bnk-gopackage"

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

func buyGlassesAtSevenEleven() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อแว่น : ที่เซเว่น : เสร็จแล้ว")
}
func buyWatchAtCentral() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อนาฬิกา : ที่เซ็นทรัล : เสร็จแล้ว")
}
func buyFruitAtSiamParagon() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อผลไม้ : ที่สยามพารากอน : เสร็จแล้ว")
}
func buyCarAtToyota() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อรถ : ที่ศูนย์โตโยต้า : เสร็จแล้ว")
}

func main() {
	fmt.Print(bnkgopackage.GetFullNameCherprang())

	start := time.Now()
	go buyGlassesAtSevenEleven()
	go buyWatchAtCentral()
	buyFruitAtSiamParagon()
	buyCarAtToyota()
	fmt.Println("ใช้เวลาในการ Run ทั้งสิ้น : ", time.Since(start), " วินาที")

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
