package main

type StudentGr struct {
	Id       int
	Number   string
	Year     int
	Students []Student
}
type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

func (st *StudentGr) avr() float64 {
	sum := 0
	for i := range st.Students {
		sum += len(st.Students[i].Rating)
	}
	return float64(sum) / float64(len(st.Students))
}

//func main() {
//	file, _ := os.Open("JSON/text.json")
//	var inp StudentGr
//	data, err := io.ReadAll(file)
//	if err != nil {
//		panic("aaa")
//	}
//	json.Unmarshal(data, &inp)
//	out := map[string]float64{
//		"Average": inp.avr(), // looking for an average number of marks
//	}
//	res, _ := json.MarshalIndent(out, "", "\t")
//	writer := bufio.NewWriter(os.Stdout)
//	writer.Write(res)
//	writer.Flush()
//
//}
