package main

import (
	"strconv"
	"net/http"
)

func numIdenticalPairs(nums []int) int {
    var sum = 0
    for i:=0; i <len(nums); i++{
        for j:=i; j<len(nums); j++{
            if(nums[i] == nums[j] && i < j){
                sum++
            }
        }
    }
    return sum
}

func handler(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	r.ParseForm()
	nums := [6] int {1, 2, 3, 1, 1, 3}
	var s[] int = nums[0:6]
	result := numIdenticalPairs(s)
	w.WriteHeader(200)
	w.Write([]byte(strconv.Itoa(result)))

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	s.ListenAndServe()
}