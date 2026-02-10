package helpers

import (
	"fmt"
	"math"
	"math/big"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/leekchan/accounting"
)

func PrintHeader() {
	pc, _, _, _ := runtime.Caller(1)
	fmt.Printf("<======> %s <======>\n", runtime.FuncForPC(pc).Name())
}

func DiffTime(a, b time.Time) string {
	var age string
	locTime := b.In(a.Location())
	_, zoneOffset := locTime.Zone()
	b = locTime.Add(-time.Duration(zoneOffset) * time.Second)

	if a.After(b) {
		a, b = b, a
	}

	diff := math.RoundToEven(b.Sub(a).Hours())
	day := int64(diff / 24)
	hour := int(diff) % 24

	if day == 0 {
		age = fmt.Sprintf("%v hours", hour)
	} else {
		age = fmt.Sprintf("%v days %v hours", day, hour)
	}

	return age
}

func ConvertStringToDate(s, layoutISO string) (time.Time, error) {
	t, err := time.Parse(layoutISO, s)
	if err != nil {
		return t, err
	}
	return t, nil
}

func CheckArray(data string, arrayCheck []string) bool {
	for _, v := range arrayCheck {
		if v == data {
			return true
		}
	}
	return false
}

func GetCurrency(current float64) string {
	ac := accounting.Accounting{Symbol: "Rp. ", Precision: 2, Thousand: ".", Decimal: ","}
	data := big.NewFloat(current)
	return ac.FormatMoneyBigFloat(data)
}

func ArrayStringToArrayInt(data []string) ([]int, error) {
	var result []int
	for _, v := range data {
		i, err := strconv.Atoi(v)
		if err != nil {
			return result, err
		}
		result = append(result, i)
	}
	return result, nil
}

func ErrorValidator(err error) string {
	var errors string
	for _, err := range err.(validator.ValidationErrors) {
		var message string
		switch err.Tag() {
		case "required":
			message = "Tidak Boleh Kosong"
		case "numeric":
			message = "Harus Berupa Angka"
		case "ne":
			message = "Karakter Tidak Diperbolehkan"
		}
		errors += fmt.Sprintf("%s %s, ", err.StructField(), message)
	}
	return errors
}

func UniqueArray(data []string) []string {
	m := make(map[string]bool)
	for _, v := range data {
		m[v] = true
	}
	var result []string
	for k := range m {
		result = append(result, k)
	}
	return result
}
