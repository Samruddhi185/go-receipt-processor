package main

import (
	"log"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func getPointsFromReceipt(receipt *Receipt) int {
	points := 0
	points += getPointsFromRetailerName(receipt)
	points += getPointsFromTotal(receipt)
	points += getPointsFromItems(receipt)
	points += getPointsFromPurchaseDate(receipt)
	points += getPointsFromPurchaseTime(receipt)
	log.Println("Total points earned: ", points)
	return points
}

func getPointsFromRetailerName(receipt *Receipt) int {
	retailerName := receipt.Retailer
	points := 0
	for _, c := range retailerName {
		if unicode.IsDigit(c) || unicode.IsLetter(c) {
			points += 1
		}
	}
	return points
}

func getPointsFromTotal(receipt *Receipt) int {
	total, _ := strconv.ParseFloat(receipt.Total, 32)
	points := 0
	if total == float64(int32(total)) {
		points += 50
	}

	if total == float64(int32(total)) ||
		(total+0.25) == float64(int32(total+0.25)) ||
		(total+0.50) == float64(int32(total+0.50)) ||
		(total+0.75) == float64(int32(total+0.75)) {
		points += 25
	}
	return points
}

func getPointsFromItems(receipt *Receipt) int {
	points := 0
	points += 5 * (len(receipt.Items) / 2)
	for _, item := range receipt.Items {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDescription)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 32)
			points += int(math.Ceil(price * 0.2))
		}
	}
	return points
}

func getPointsFromPurchaseDate(receipt *Receipt) int {
	points := 0
	purchaseDate, _ := strconv.Atoi(strings.Split(receipt.PurchaseDate, "-")[2])
	if purchaseDate%2 == 1 {
		points += 6
	}
	return points
}

func getPointsFromPurchaseTime(receipt *Receipt) int {
	points := 0
	splitTime := strings.Split(receipt.PurchaseTime, ":")
	hour, _ := strconv.Atoi(splitTime[0])
	mins, _ := strconv.Atoi(splitTime[1])
	if (hour == 14 && mins > 0) || (hour == 15) {
		points += 10
	}
	return points
}
