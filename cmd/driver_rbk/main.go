package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	var (
		hex1       = "948E48C23D0E6B656F570BC7CDD2DCBC"
		hex2       = "FE5E04258C1871AF247077FE3EB9DDFA"
		terminalID = "EP104059"
		currency   = "KZT"
	)

	key, err := calculateXOR(hex1, hex2)
	if err != nil {
		log.Println(1, err)
		return
	}

	body, xErr := newRefundRequest(terminalID, currency).body(key)
	if xErr != nil {
		log.Println(2, err)
		return
	}

	log.Printf("%s", body)
}

func calculateXOR(hex1, hex2 string) (string, error) {
	num1, err := hex.DecodeString(hex1)
	if err != nil {
		return "", fmt.Errorf("invalid component 1: %s. Err: %s", hex1, err.Error())
	}

	num2, err := hex.DecodeString(hex2)
	if err != nil {
		return "", fmt.Errorf("invalid component 2: %s. Err: %s", hex2, err.Error())
	}

	resultBytes := make([]byte, len(num1))
	for i := 0; i < len(num1); i++ {
		resultBytes[i] = num1[i] ^ num2[i]
	}

	return hex.EncodeToString(resultBytes), nil
}

type refundRequest struct {
	Amount    string `json:"amount"`
	Currency  string `json:"currency"`
	Rrn       string `json:"rrn"`
	IntRef    string `json:"intRef"`
	TrType    string `json:"trtype"`
	Terminal  string `json:"terminal"`
	Timestamp string `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Sign      string `json:"signature"`
}

func newRefundRequest(tid, currency string) *refundRequest {
	timestamp := time.Now().UTC().Format("20060102150405")

	return &refundRequest{
		Amount:    "15000",
		Currency:  currency,
		Rrn:       "426983499508",
		IntRef:    "5D1EC9EDC16F5475",
		TrType:    "174",
		Terminal:  tid,
		Timestamp: timestamp,
		Nonce:     generateNonce(),
	}
}

func generateNonce() string {
	minLength := 16
	maxLength := 32
	length := rand.Intn(maxLength-minLength+1) + minLength //nolint:gosec // ..

	nonce := ""
	characters := "0123456789ABCDEF"

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(characters)) //nolint:gosec // ..
		nonce += string(characters[randomIndex])
	}

	return nonce
}

func (r *refundRequest) body(key string) ([]byte, error) {
	r.Sign = createSign(key, r.Terminal, r.TrType, r.Amount, "")
	return json.Marshal(r)
}

func createSign(key, terminal, trType, amount, desc string) string {
	var sb strings.Builder
	sb.WriteString(getSize(terminal))
	sb.WriteString(terminal)
	sb.WriteString(getSize(trType))
	sb.WriteString(trType)
	sb.WriteString(getSize(amount))
	sb.WriteString(amount)
	if desc != "" {
		sb.WriteString(getSize(desc))
		sb.WriteString(desc)
	} else {
		sb.WriteString("-")
	}

	ds, _ := hex.DecodeString(key)
	mac := hmac.New(sha1.New, ds)
	mac.Write([]byte(sb.String()))
	return hex.EncodeToString(mac.Sum(nil))
}

func getSize(s string) string {
	return strconv.Itoa(len(s))
}
