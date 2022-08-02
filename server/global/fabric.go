package global

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

var (
	SDK              *fabsdk.FabricSDK
	LEDGER_CLIENT    *ledger.Client
	CHANNEL_CLIENT   *channel.Client
	CHANNEL_PROVIDER context.ChannelProvider
	GW               *gateway.Gateway
	NETWORK          *gateway.Network
	CONTRACT         *gateway.Contract
	CLIENT_PROVIDER  context.ClientProvider
	RC               *resmgmt.Client
)
