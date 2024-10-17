package config

// # *下面是Hmac256签名方式的变量，SIGN是“256”代表Hmac256签名方式，“25519”代表Ed25519签名方式。
// # * API_KEY和SECRET_KEY是Hmac256方式需要的公钥和私钥。PUBLIC_KEY和PRIVATE_KEY是Ed25519的公钥和私钥。
// # *
// # * The following is the variable of the Hmac256 signature mode.
// # * SIGN indicates the Hmac256 signature mode and 25519 indicates the Ed25519 signature mode.
// # * API_KEY and SECRET_KEY are the public and private keys required for Hmac256 mode.
// # * PUBLIC_KEY and PRIVATE_KEY are the public and private keys of Ed25519.

var Host = "dawn-broker-pro.global-base.tc-jp1.huobiapps.com"
var AccountId = "31252055"
var SubUid int64 = 5678
var SubUids string = "5678"


var AccessKey = "2c47b6d2-ed2htwf5tf-eab8ce25-e441a"
var SecretKey = "e44fd253-1f451fc8-d5076472-fd4a0"
var Sign = "256"

var PublicKey = "65078cb9-bvrge3rf7j-5b59ba25-22845"
var PrivateKey = `-----BEGIN PRIVATE KEY-----
MC4CAQAwBQYDK2VwBCIEIM+xKGLUEX92/bYpWD/U+nAOGMwm/I+TAOl34xZGp0+6
-----END PRIVATE KEY-----`
// var Sign = "25519"


