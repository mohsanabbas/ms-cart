package cart

// User userData structue
type User struct {
	Credential struct {
		Personid   int         `json:"personId" bson:"personId"`
		Userid     int         `json:"userId" bson:"userId"`
		VendorName string      `json:"name" bson:"name"`
		Email      string      `json:"email" bson:"email"`
		Cpf        interface{} `json:"cpf" bson:"cpf"`
		Branchid   int         `json:"branchId" bson:"branchId"`
		Agentsign  string      `json:"agentSign" bson:"agentSign"`
		User       string      `json:"user" bson:"user"`
		Usertype   string      `json:"userType" bson:"userType"`
	} `json:"credential" bson:"credential"`
	Systems []interface{} `json:"systems" bson:"systems"`
	Iat     int64         `json:"iat" bson:"iat"`
}
