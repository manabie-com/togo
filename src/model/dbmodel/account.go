package dbmodel

// Account structure account in DB
type Account struct {
	Account           string `bson:"account" json:"account"`
	Password          string `bson:"password" json:"password"`
	Group             Group  `bson:"group" json:"group"`
	CounterTaskPerDay int32  `bson:"counterTaskPerDay" json:"counterTaskPerDay"`
	Task              []Task `bson:"task" json:"task"`
}

// Group structure in db
type Group struct {
	GroupName         string `bson:"groupName" json:"groupName"`
	MaximumTaskPerDay int32  `bson:"MaximumTaskPerDay" json:"MaximumTaskPerDay"`
}

var GroupDemo []Group

func init() {
	accountant := Group{GroupName: "accountant", MaximumTaskPerDay: 2}
	security := Group{GroupName: "security", MaximumTaskPerDay: 3}
	staff := Group{GroupName: "staff", MaximumTaskPerDay: 4}
	deputy := Group{GroupName: "deputy", MaximumTaskPerDay: 5}
	manager := Group{GroupName: "manager", MaximumTaskPerDay: 6}
	GroupDemo = []Group{accountant, security, staff, deputy, manager}
}
