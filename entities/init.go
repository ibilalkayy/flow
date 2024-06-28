package entities

type AuthVariables struct {
	Username    string
	Gmail       string
	AppPassword string
}

type DatabaseVariables struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type BlockchainVariables struct {
	PrivateKey    string
	AlchemyApiURL string
}
