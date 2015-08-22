package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data := UploadData("data.csv")

	var x_vals, y_vals []float64

	for _, each := range data[1:] {
		x, _ := strconv.ParseFloat(each[1], 64)
		y, _ := strconv.ParseFloat(each[2], 64)

		x_vals = append(x_vals, x)
		y_vals = append(y_vals, y)
	}

	x_avg := mean(x_vals)
	x_devs := devs(x_vals, x_avg)
	x_sqrs := squares(x_devs)
	for _, num := range x_sqrs {
		// fmt.Printf("x: %0.1f, x_avg: %0.1f, (x - x_avg): %0.1f\n", num, x_avg, sums(y_vals))
		fmt.Println(num)
	}

	y_avg := mean(y_vals)
	y_devs := devs(y_vals, y_avg)
	y_sqrs := squares(y_devs)
	for _, num := range y_sqrs {
		fmt.Println(num)
	}

	x_prods := products(x_devs, y_devs)
	for _, num := range x_prods {
		fmt.Println(num)
	}

	fmt.Printf("x_sum: %0.1f, y_sum: %0.1f\n", sums(x_vals), sums(y_vals))
	fmt.Printf("x_avg: %0.1f, y_avg: %0.1f\n", mean(x_vals), mean(y_vals))
	fmt.Printf("x_sqrs_sum: %0.1f, y_sqrs_sum: %0.1f\n", sums(x_sqrs), sums(y_sqrs))
	fmt.Printf("x_prods_sum: %0.1f\n", sums(products(x_devs, y_devs)))

	fmt.Printf("x_coef: %0.3f\n", sums(products(x_devs, y_devs))/sums(x_sqrs))
	fmt.Printf("slope: %0.3f\n", y_avg-((sums(products(x_devs, y_devs))/sums(x_sqrs))*x_avg))

}

func mean(vals []float64) float64 {
	total := sums(vals)
	return total / float64(len(vals))
}

func sums(vals []float64) (total float64) {
	for _, num := range vals {
		total += num
	}
	return
}

func devs(vals []float64, target float64) (dvs []float64) {
	for _, num := range vals {
		dvs = append(dvs, num-target)
	}
	return
}

func squares(vals []float64) (sqrs []float64) {
	for _, num := range vals {
		sqrs = append(sqrs, num*num)
	}
	return
}

func products(x []float64, y []float64) (prods []float64) {
	if len(x) != len(y) {
		os.Exit(1)
	}

	for i := 0; i < len(x); i++ {
		prods = append(prods, x[i]*y[i])
	}
	return
}

func determinate(y_vals []float64) {
	N := len(y_vals)
}

func UploadData(fn string) (raw [][]string) {
	file, err := os.Open(fn)
	if err != nil {
		fmt.Println(err)
		return raw
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // see the Reader struct information below
	raw, err = reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}
