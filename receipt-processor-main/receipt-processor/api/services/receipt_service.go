package services

import (
	"fmt"
	"math"
	"receipt-processor/api/models"
	"receipt-processor/api/utils"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var validDecimal = regexp.MustCompile(`^\d+\.\d{2}$`)
var receipts = make(map[string]models.Receipt)
var receiptPoints = make(map[string]int)
var mu sync.Mutex

func ValidateReceipt(receipt *models.Receipt) bool {
	if receipt.Retailer == "" || receipt.PurchaseDate == "" || receipt.PurchaseTime == "" || len(receipt.Items) == 0 || receipt.Total == "" {
		return false
	}

	if !validDecimal.MatchString(receipt.Total) {
		return false
	}

	return true
}

func ProcessReceipt(receipt *models.Receipt) string {
	receiptID := utils.GenerateUUID()

	mu.Lock()
	receipts[receiptID] = *receipt
	mu.Unlock()

	points := CalculatePoints(receipt)
	mu.Lock()
	receiptPoints[receiptID] = points
	mu.Unlock()

	return receiptID
}

func GetReceiptPoints(receiptID string) int {
	mu.Lock()
	points, exists := receiptPoints[receiptID]
	mu.Unlock()

	if !exists {
		return -1
	}

	return points
}

func CalculatePoints(receipt *models.Receipt) int {
	points := 0

	// One point for every alphanumeric character in the retailer name
	alphanumericRetailerName := regexp.MustCompile(`[A-Za-z0-9]+`).FindAllString(receipt.Retailer, -1)
	points += len(strings.Join(alphanumericRetailerName, ""))

	// 50 points if the total is a round dollar amount with no cents
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if math.Mod(total, 1.0) == 0.0 {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25
	if math.Mod(total, 0.25) == 0.0 {
		points += 25
	}

	//5 points for every two items on the receipt
	numItems := len(receipt.Items)
	points += (numItems / 2) * 5

	// If the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer.
	for _, item := range receipt.Items {
		description := strings.TrimSpace(item.ShortDescription)
		if len(description)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// 6 points if the day in the purchase date is odd
	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	fmt.Print(time.Date(0, 0, 0, 14, 0, 0, 0, time.UTC))
	if purchaseTime.After(time.Date(purchaseTime.Year(), purchaseTime.Month(), purchaseTime.Day(), 14, 0, 0, 0, time.UTC)) && purchaseTime.Before(time.Date(purchaseTime.Year(), purchaseTime.Month(), purchaseTime.Day(), 16, 0, 0, 0, time.UTC)) {
		points += 10
	}

	return points
}
