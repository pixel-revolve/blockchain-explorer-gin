package utils

import (
	"fmt"
	"gin/global"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"log"
)

func SubmitTransaction(function string, params ...string) []byte {
	result, err := global.CONTRACT.SubmitTransaction(function, params...)
	if err != nil {
		returnErr := fmt.Sprintf("Failed to Submit Transaction: %s\n", err)
		return []byte(returnErr)
	}
	return result
}

func CreateLedgerClient(channelID string, username string) (*ledger.Client, error) {
	// 根据channelID获取到通道。
	channelProvider := global.SDK.ChannelContext(channelID, fabsdk.WithUser(username))
	ledgerClient, err := ledger.New(channelProvider)
	if err != nil {
		log.Fatalln("Failed to create new ledgerClient: ", err)
		return nil, err
	}
	return ledgerClient, nil
}
