package main

//func main() {
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, "Main page")
//
//	})
//	// will return the square of the circle if input radius to the URL parameter
//	http.HandleFunc("/rad", func(w http.ResponseWriter, r *http.Request) {
//		s := string(r.URL.Query().Get("radius"))
//		i, _ := strconv.Atoi(s)
//		fl := float64(i*i) * math.Pi
//		str := strconv.FormatFloat(fl, 'f', -1, 64)
//		fmt.Fprintf(w, str)
//
//	})
//	//Simple server
//	http.ListenAndServe(":8080", nil)
//
//}
