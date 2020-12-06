package define

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

var (
	// ContextKeyDeleteCaller var
	ContextKeyUserID = contextKey("user_id")
)
