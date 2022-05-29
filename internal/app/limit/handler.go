package limit

import "context"

type (
	Service interface {
		GetLimit(ctx context.Context, req *GetLimitReq) (*Limit, error)
	}
)

// TODO: implement later
