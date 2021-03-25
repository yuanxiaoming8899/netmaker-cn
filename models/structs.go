package models

import jwt "github.com/dgrijalva/jwt-go"

type AuthParams struct {
    MacAddress    string `json:"macaddress"`
    Password string `json:"password"`
}

type User struct {
    UserName string `json:"username" bson:"username" validate:username_valid,username_unique,min=3`
    Password string `json:"password" bson:"password" validate:password_check`
    IsAdmin bool `json:"isadmin" bson:"isadmin"`
}

type UserAuthParams struct {
    UserName    string `json:"username"`
    Password string `json:"password"`
}

type UserClaims struct {
    IsAdmin bool
    UserName string
    jwt.StandardClaims
}

type SuccessfulUserLoginResponse struct {
    UserName     string
    AuthToken string
}

// Claims is  a struct that will be encoded to a JWT.
// jwt.StandardClaims is an embedded type to provide expiry time
type Claims struct {
    Group string
    MacAddress string
    jwt.StandardClaims
}

// SuccessfulLoginResponse is struct to send the request response
type SuccessfulLoginResponse struct {
    MacAddress     string
    AuthToken string
}

type ErrorResponse struct {
    Code    int
    Message string
}

type NodeAuth struct {
    Group    string
    Password string
    MacAddress string
}

// SuccessResponse is struct for sending error message with code.
type SuccessResponse struct {
    Code     int
    Message  string
    Response interface{}
}

type AccessKey struct {
    Name string `json:"name" bson:"name"`
    Value string `json:"value" bson:"value"`
    Uses int `json:"uses" bson:"uses"`
}

type DisplayKey struct {
    Name string `json:"name" bson:"name"`
    Uses int `json:"uses" bson:"uses"`
}

type CheckInResponse struct{
    Success bool `json:"success" bson:"success"`
    NeedPeerUpdate bool `json:"needpeerupdate" bson:"needpeerupdate"`
    NeedConfigUpdate bool `json:"needconfigupdate" bson:"needconfigupdate"`
    NodeMessage string `json:"nodemessage" bson:"nodemessage"`
    IsPending bool `json:"ispending" bson:"ispending"`
}

type PeersResponse struct {
    PublicKey string `json:"publickey" bson:"publickey"`
    Endpoint string `json:"endpoint" bson:"endpoint"`
    Address string `json:"address" bson:"address"`
    LocalAddress string `json:"localaddress" bson:"localaddress"`
    ListenPort int32 `json:"listenport" bson:"listenport"`
    KeepAlive int32 `json:"persistentkeepalive" bson:"persistentkeepalive"`
}
