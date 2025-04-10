package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateTransactionRedeemCode() string {
	generateCode := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("REDEEM-%d-%04d", time.Now().Unix(), generateCode.Intn(10000))
}