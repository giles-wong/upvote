package wedoctor

const (
	ClientId      string = "d4d6ff4c5174edf7f7d5e9d3ee64b23b92a97e60"
	SignatureSalt string = "6e1cc053f8a113831d124a5a788835b673deb06a1af5ed9008f531299b7df404e4ed8ce9294501858ea76950fe540787c2e66996ea4fb0149eec80ec1bec4112s"
	Domain        string = "https://gateway.guahao.cn"
)

type ApiGatewayProperties struct {
	clientId      string
	clientSecret  string
	scope         string
	signatureSalt string
	serverUrl     string
	grantType     string
}

func buildSignature() {

}

func getAccessToken() {

}

// Get 发送Get 请求
func Get(path string, param interface{}, headers interface{}) {

}

// Post 发送Get 请求
func Post(path string, param interface{}, headers interface{}) {

}
