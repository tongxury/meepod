package enum

import "time"

const PaymentTimeout = 3 * time.Minute
const PaymentTimeoutMinute = 15

const MinProxyRewardRate = float64(0.01)
const MaxProxyRewardRate = float64(0.07)

//const Clean = int(2)
//const NotClean = int(1)

const DefaultStoreIcon = "default_store_icon.png"
const DefaultUserIcon = "default_user_icon.png"

const CoStoreRewardRate = float64(0.055)
const CoStoreRewardMax = float64(50)
const ShowConStore = false

var SecretSignKey = []byte("7c94CBNXSyGWf6Ir")
