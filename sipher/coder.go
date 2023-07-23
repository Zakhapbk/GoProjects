package main

import (
	"bufio"
	"io"
	"os"
)

func siph(n int, str string) string {
	nstr := []byte(str)
	for i := range nstr {
		if nstr[i] == 13 || nstr[i] == 10 {
			break
		}
		nstr[i] = nstr[i] + byte(n%25)
	}
	str = string(nstr)
	return str
}

func desiph(n int, str string) string {
	nstr := []byte(str)
	for i := range nstr {
		if nstr[i] == 13 || nstr[i] == 10 {
			break
		}
		nstr[i] = nstr[i] - byte(n%25)
	}
	str = string(nstr)
	return str
}

func siphFile() {
	file, err := os.Open("sipher/pswd.txt")
	if err != nil {
		panic("No file")
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	nfile, err := os.Create("sipher/sipher.txt")
	if err != nil {
		panic("No file created")
	}
	str, err := "", nil
	w := bufio.NewWriter(nfile)
	for true {
		str, err = rd.ReadString('\n')
		str = siph(3, str)
		_, _ = w.WriteString(str)
		if err == io.EOF {
			w.Flush()
			break
		}
	}
}

func desiphFile() {
	file, err := os.Open("sipher/sipher.txt")
	if err != nil {
		panic("No file")
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	nfile, err := os.Create("sipher/desipher.txt")
	if err != nil {
		panic("No file created")
	}
	str, err := "", nil
	w := bufio.NewWriter(nfile)
	for true {
		str, err = rd.ReadString('\n')
		str = desiph(3, str)
		_, _ = w.WriteString(str)
		if err == io.EOF {
			w.Flush()
			break
		}
	}
}

//func main() {
//	siphFile()   //encodes the file pswd.txt to sipher.txt (using Caesar sipher)
//	desiphFile() // decodes sipher.txt to desipher.txt
//}
