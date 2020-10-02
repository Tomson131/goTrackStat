package main

import "os"

func GetLine(l string) {
	println(l)
}
func Contains(ar []BarCode, stat StatusBase) bool {
	for _, n := range ar {
		if stat.Barcode == n.Barcode {
			return true
		}
	}
	return false
}
func SetContains(ar []BarCode, stat StatusBase) []BarCode {
	for i, n := range ar {
		if stat.Barcode == n.Barcode {
			ar[i].List = append(n.List, stat)
		}
	}
	return ar
}

/*func AddStatar(ar []StatusBase, s StatusBase) []StatusBase {
	ar = append(ar, s)
	return ar
}*/
func AddBarar(ar []BarCode, s StatusBase) []BarCode {
	ar = append(ar, BarCode{Barcode: s.Barcode, List: []StatusBase{s}})
	return ar
}

func WriteToFile(filename string, msg string) {
	file, err := os.Create(filename)
	if err != nil {
		println("Unable to create file:", err)
	}
	defer file.Close()
	file.WriteString(msg)
}
