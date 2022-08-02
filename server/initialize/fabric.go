package initialize

import (
	"fmt"
	"gin/global"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"log"
)

const (
	configPath  = "./e2e.yaml"
	channelID   = "mychannel"
	username    = "Admin"
	chaincodeId = "basic"
	org         = "Org1"
)

var (
	sdk             = &global.SDK
	ledgerClient    = &global.LEDGER_CLIENT
	channelClient   = &global.CHANNEL_CLIENT
	channelProvider = &global.CHANNEL_PROVIDER
	clientProvider  = &global.CLIENT_PROVIDER
	gw              = &global.GW
	network         = &global.NETWORK
	contract        = &global.CONTRACT
	rc              = &global.RC
)

// 下次编写的时候尽量把initialize和global给写到一块去
func sdkInit() {
	if *sdk != nil {
		return
	}
	var err error
	// 通过config包方法从e2e.yaml加载配置。
	configProvider := config.FromFile(configPath)
	// 通过配置初始化sdk
	*sdk, err = fabsdk.New(configProvider)
	if err != nil {
		log.Fatalf("Failed to create new SDK: %s\n", err)
		return
	}
	// 根据channelID获取到通道。
	*channelProvider = (*sdk).ChannelContext(channelID, fabsdk.WithUser(username))
	*clientProvider = (*sdk).Context(fabsdk.WithUser(username), fabsdk.WithOrg(org))
	return
}

// 初始化资源管理客户端
func rcInit() {
	if *rc != nil {
		return
	}
	var err error
	*rc, err = resmgmt.New(*clientProvider)
	if err != nil {
		log.Panicf("failed to create resource client: %s", err)
	}
}

// 在这个例子中ledger用于查询区块链的基本信息，例如区块数量等。
func ledgerClientInit() {
	if *ledgerClient != nil {
		return
	}
	var err error
	*ledgerClient, err = ledger.New(*channelProvider)
	if err != nil {
		log.Fatalln("Failed to create new ledgerClient: ", err)
		return
	}
}

// 获取channel客户吨
func channelClientInit() {
	if *channelClient != nil {
		return
	}
	var err error
	*channelClient, err = channel.New(*channelProvider)
	if err != nil {
		log.Fatalln("Failed to create new channelClient: ", err)
		return
	}
}

// gateway还是用于连接fabric网络，然后调用合约
func gatewayInit() {
	if *gw != nil {
		return
	}
	var err error
	*gw, err = gateway.Connect(gateway.WithSDK(*sdk), gateway.WithUser("Admin"))
	if err != nil {
		fmt.Printf("Failed to create gateway: %s", err)
		return
	}
}

func networkInit() {
	if *network != nil {
		return
	}
	var err error
	*network, err = (*gw).GetNetwork(channelID)
	if err != nil {
		fmt.Printf("Failed to connect network %s", err)
		return
	}
}

func contractInit() {
	if *contract != nil {
		return
	}
	*contract = (*network).GetContract(chaincodeId)
}

func GetContract() *gateway.Contract {
	return *contract
}

func FabricInit() {
	sdkInit()
	rcInit()
	ledgerClientInit()
	channelClientInit()
	gatewayInit()
	networkInit()
	contractInit()
	fmt.Println("初始化完成")
}
