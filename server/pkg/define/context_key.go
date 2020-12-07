package define

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

var (
	ContextKeyAuthorization = contextKey("Authorization")
	ContextKeyUserID = contextKey("user_id")
)
