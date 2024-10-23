package config

// # *下面是Hmac256签名方式的变量，SIGN是“256”代表Hmac256签名方式，“25519”代表Ed25519签名方式。
// # * API_KEY和SECRET_KEY是Hmac256方式需要的公钥和私钥。PUBLIC_KEY和PRIVATE_KEY是Ed25519的公钥和私钥。
// # *
// # * The following is the variable of the Hmac256 signature mode.
// # * SIGN indicates the Hmac256 signature mode and 25519 indicates the Ed25519 signature mode.
// # * API_KEY and SECRET_KEY are the public and private keys required for Hmac256 mode.
// # * PUBLIC_KEY and PRIVATE_KEY are the public and private keys of Ed25519.

var Host = ""
var AccountId = "31252055"
var SubUid int64 = 5678
var SubUids string = "5678"

var AccessKey = ""
var SecretKey = ""
var Sign = "256"

var PublicKey = ""
var PrivateKey = ``

// var Sign = "25519"
