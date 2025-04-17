package utils

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ShashankRKumar/receipt-processor/models"
)

func CalculatePoints(receipt models.Receipt) int {
	points := 0
	reg := regexp.MustCompile(`[a-zA-Z0-9]`)
	points += len(reg.FindAllString(receipt.Retailer, -1))

	if isRoundDollar(receipt.Total) {
		points += 50
	}
	if isMultipleOfQuarter(receipt.Total) {
		points += 25
	}

	points += (len(receipt.Items) / 2) * 5

	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	totalFloat, _ := strconv.ParseFloat(receipt.Total, 64)
	if totalFloat > 10.0 {
		points += 5
	}

	if isPurchaseDayOdd(receipt.PurchaseDate) {
		points += 6
	}

	if isBetweenTwoAndFour(receipt.PurchaseTime) {
		points += 10
	}

	return points
}

func isRoundDollar(total string) bool {
	f, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return false
	}
	return f == float64(int64(f))
}

func isMultipleOfQuarter(total string) bool {
	f, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return false
	}
	return math.Mod(f, 0.25) == 0
}

func isPurchaseDayOdd(date string) bool {
	parsed, err := time.Parse("2006-01-02", date)
	if err != nil {
		return false
	}
	return parsed.Day()%2 != 0
}

func isBetweenTwoAndFour(timeStr string) bool {
	t, err := time.Parse("15:04", timeStr)
	if err != nil {
		return false
	}
	hour := t.Hour()
	minute := t.Minute()
	return hour == 14 || (hour == 15 && minute == 0)
}
