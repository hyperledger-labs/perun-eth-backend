package liquiditypool

import "github.com/ethereum/go-ethereum/common"

// Event signature topics for LiquidityPool reconciliation subscriptions.
var (
	LiquidityPoolChannelFundedTopic   = common.HexToHash("0x921e723c5ab46e7fad6ecc352c41767d15f80a4d322198355af1f032581da3ea")
	LiquidityPoolChannelSettledTopic  = common.HexToHash("0xbb7a5a625951cb72f1e443bc7a3cd146dcdc284df4a5ed440993e5df803d2bc7")
	LiquidityPoolChannelExpiredTopic  = common.HexToHash("0xd2c7b3913e8b91b3448548f3fc1965a5455e0fa256af80705cab84affce61fa6")
	LiquidityPoolBondPostedTopic      = common.HexToHash("0xc4e1648d08804ab9c896cf4e8349ac6b9d4849cbeffe50bb73a3b9d7970a8c54")
	LiquidityPoolBondWithdrawnTopic   = common.HexToHash("0xf0566a10f495c405474bbbfa43e7d77dffab0bdbf3240d1880b2b06249685409")
	LiquidityPoolOperatorUpdatedTopic = common.HexToHash("0xfbe5b6cbafb274f445d7fed869dc77a838d8243a22c460de156560e8857cad03")
)
