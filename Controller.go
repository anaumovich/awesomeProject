package Main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)
var i = 0

var Store = make(map[int]Product)

func main() {
	handler := createHandler(getRoutes())

	http.HandleFunc("/", http.HandlerFunc(handler))
	http.ListenAndServe(":8080", nil)

}

func getRoutes() (m map[string]func(w http.ResponseWriter, r *http.Request)) {
	m = make(map[string]func(w http.ResponseWriter, r *http.Request))
	m["GET/"] = listAdd
	m["POST/list"] = GetList
	m["POST/redirect"] = ReturnToHome
	m["GET/edit"] = EditData // regexp
	m["POST/addProduct"] = AddData
	return m
}

func createHandler(m map[string]func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := strings.Join([]string{r.Method, r.URL.Path}, "")

		fmt.Println(key)

		if result, ok := m[key]; ok == true {
			result(w, r)
		} else {
			w.WriteHeader(404)
		}
	}
}

func Add(name string, count int64, price float64, i int, x map[int]Product) {
	x[i] = Product{name, count, price}
}

func listAdd(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(startPage()))
}

func GetList(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("First")
	count, _ := strconv.ParseInt(r.FormValue("Second"), 10, 64)
	price, _ := strconv.ParseFloat(r.FormValue("Third"), 64)
	//fmt.Println(count)
	Add(name, count, price, i, Store)
	i++
	//fmt.Println(Store[1].name)
	w.Header().Set("Location", "http://localhost:8080/add")
	_, _ = w.Write([]byte(newPage(addString())))
}

func ReturnToHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://localhost:8080")
	w.WriteHeader(302)
}

func EditData(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(EditProduct()))
	//id, _ = strconv.Atoi(r.URL.Fragment(`edit?product_id=` + strconv.Itoa(a)`))
}

func AddData(w http.ResponseWriter, r *http.Request) {
	editName := r.FormValue("One")
	editCount, _ := strconv.ParseInt(r.FormValue("Two"), 10, 64)
	editPrice, _ := strconv.ParseFloat(r.FormValue("Three"), 64)
	ChangeProduct( w , r , editName, editCount, editPrice)

	_, _ = w.Write([]byte(newPage(addString())))
}

func addString() string {

	b := ""

	for a := 1; a < i; a++ {
		arr := make([]string, 8)
		arr[0] = `<tr><td>`
		arr[1] = string(Store[a].name)
		arr[2] = `</td><td>`
		arr[3] = strconv.FormatInt(Store[a].count, 10)
		arr[4] = `</td><td>`
		arr[5] = strconv.FormatFloat(Store[a].price, 'f', 0, 64)
		arr[6] = `</td><td><a href="http://localhost:8080/edit?product_id=` + strconv.Itoa(a) + `"><button>Изменить</button></a></td>`
		arr[7] = `</td></tr>`
		b += strings.Join(arr, "")
	}
	return b
}

func ChangeProduct( w http.ResponseWriter, r *http.Request, editName string, editCount int64, editPrice float64) {
	Store[strconv.Atoi(r.URL.Fragment(`edit?product_id=` + strconv.Itoa(a)``))] = Product{editName, editCount, editPrice}
}