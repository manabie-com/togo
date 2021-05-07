package jwt_service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"manabie-com/togo/global"
	"manabie-com/togo/util"
	"time"
)

const JWTPrivateRSKey = "-----BEGIN RSA PRIVATE KEY-----\nMIICXAIBAAKBgQCgKk5JPoD41rAF2+3mST9dOGdpHNW7YcSHMM2q2+EvdRF1KcL9\nu8Pw4WCXbfCmHbpGF5z7PsAkEJHe1Yc75rR9/OG7mUgnok9Tt1/x9MVUCYZdJky9\nkOxCdfm1sTpnFUGmemBv5Dg0ekX3AHAc1hU6lGHVM/0wCnxiBYNLt434OQIDAQAB\nAoGAUDGejRHcpbto2yYpUbqvHU/Wh0zGv1HOgxougDQj5g0pto44ca8IBp3yLSAA\n9EvCLsI3+ZyLvAMH5pjnY1i6WeeQH5u4enw8oWbGz/Vh8cdRG9oGgWjNbHVqJJyJ\nATujon3bfoQAOhcyJalWc7biEtCTAjblut/Rj+baRWQlImECQQDXdIqW68PgHw92\nWI1MsULD3g4gMd0E+yUlhX/JUPbVF/xvEn7B1fu2RJuPtZco7ITjQYWhAJMJ4gBM\niQGcOoLdAkEAvk4xzmLlAk7973xFLv/GPhlAGTNM5mLx4mKQMhuKtV3PCwYSN4El\neBgph7IKnrZTZyc/gPM01vB7VvEs5P3vDQJBAMUHUmXJnQqr3NwBBtaHk+LCgnB2\nqQQRF1tExiM340Hj+XkplLl2EgYQn6HAEkfeY3ffR3CAsfZrspJLCCnyaBECQH3p\n+uZVZLTMUxP7o0LflOlNh62k1cKxwN1K3aFpu7MYqH7gu3jiCEqXohLYaFJuzGw5\n+bh2MoXsg48Y791rbpkCQFfcpAASgD/3620V6bWhd1muF2lkxMgJ9i74ybaoZ4AH\nCfUXiYX6mRWh1Jo2kGgpSVKi62PptP4zKrMsELmeHnQ=\n-----END RSA PRIVATE KEY-----"
const JWTPublicRSKey = "-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCgKk5JPoD41rAF2+3mST9dOGdp\nHNW7YcSHMM2q2+EvdRF1KcL9u8Pw4WCXbfCmHbpGF5z7PsAkEJHe1Yc75rR9/OG7\nmUgnok9Tt1/x9MVUCYZdJky9kOxCdfm1sTpnFUGmemBv5Dg0ekX3AHAc1hU6lGHV\nM/0wCnxiBYNLt434OQIDAQAB\n-----END PUBLIC KEY-----"

type Claims struct {
	Id      string `json:"id"`
	MaxTodo int    `json:"max_todo"`
	jwt.StandardClaims
}

func CreateJwt(cl Claims) string {
	var claims = Claims{
		Id:      cl.Id,
		MaxTodo: cl.MaxTodo,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(global.Config.JwtValidMinute)).Unix(),
			Issuer:    "MANABIE",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	// Sign and get the complete encoded token as a string using the secret
	privateKey, _ := jwt.ParseRSAPrivateKeyFromPEM([]byte(JWTPrivateRSKey))
	tokenString, _ := token.SignedString(privateKey)

	// optional if we don't want anyone can see jwt body
	var securityToken, _ = util.AESEncrypt([]byte(global.Config.AESJWTKey), tokenString)
	return securityToken
}

func DoAuthenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		securityToken := c.Request.Header.Get("Authorization")

		var token string
		var aesDecryptErr error
		token, aesDecryptErr = util.AESDecrypt([]byte(global.Config.AESJWTKey), securityToken)
		if aesDecryptErr != nil {
			util.AbortUnauthorized(c, util.ERR_CODE_DECRYPT_JWT_TOKEN, "Can't decrypt token!")
			return
		}

		publicKey, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(JWTPublicRSKey))
		var claims Claims
		result, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (i interface{}, e error) {
			return publicKey, nil
		})

		if result != nil && result.Valid {
			c.Set("user_id", claims.Id)
			c.Set("max_todo", claims.MaxTodo)
			return
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				util.AbortUnauthorized(c, util.ERR_CODE_NOT_EVEN_A_TOKEN, "That's not even a token!")
				return
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				util.AbortUnauthorized(c, util.ERR_CODE_TOKEN_EXPIRED, "Token expired!")
				return
			} else {
				util.AbortUnauthorized(c, util.ERR_CODE_COULD_NOT_HANDLE_THIS_TOKEN, "Couldn't handle this token!")
				return
			}
		} else {
			util.AbortUnauthorized(c, util.ERR_CODE_INVALID_TOKEN, "Couldn't handle this token!")
			return
		}
	}
}
