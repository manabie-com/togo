package auth_services

// policy description
var PolicyDescription = map[string]string{
	"admin":    "All permission: 1. user management 2. configurations 3. monitoring",
	"operator": "All permissions except User management: 1. configurations, 2. monitoring",
	"monitor":  "Only Monitoring permission",
	"unassign": "Nothing can do",
}
