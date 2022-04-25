package errors

import "testing"

func TestGetError(t *testing.T) {
	type args struct {
		errContext int
		errMessage string
		errCode    string
	}

	tests := []struct {
		name string
		args
		want CustomError
	}{
		{
			name: "Request Invalid",
			args: args{
				BadRequestContext,
				BadRequestMessage,
				TodoTaskRequestInvalid,
			},
			want: CustomError{
				Code:    BadRequestContext,
				Message: BadRequestMessage,
				ErrorInfor: ErrorInfor{
					Code:        TodoTaskRequestInvalid,
					Message:     "To Do Task Request is invalid",
					Description: "Invalid or missing requested fields on to do task request body! Please check and try again!",
				},
			},
		},
		{
			name: "Exceed Task Limit",
			args: args{
				BadRequestContext,
				BadRequestMessage,
				ExceedDailyLimitRecords,
			},
			want: CustomError{
				Code:    BadRequestContext,
				Message: BadRequestMessage,
				ErrorInfor: ErrorInfor{
					Code:        ExceedDailyLimitRecords,
					Message:     "Exceed Daily Limit Records",
					Description: "This user has exceeded daily limit of to do records",
				},
			},
		},
		{
			name: "User Id Not Found",
			args: args{
				BadRequestContext,
				BadRequestMessage,
				UserIdNotFound,
			},
			want: CustomError{
				Code:    BadRequestContext,
				Message: BadRequestMessage,
				ErrorInfor: ErrorInfor{
					Code:        UserIdNotFound,
					Message:     "User Id is not existed",
					Description: "Requested User Id is not existed! Please try again!",
				},
			},
		},
		{
			name: "Unexpected Error",
			args: args{
				InternalErrorContext,
				InteralErrorMessage,
				UnexpectedError,
			},
			want: CustomError{
				Code:    InternalErrorContext,
				Message: InteralErrorMessage,
				ErrorInfor: ErrorInfor{
					Code:        UnexpectedError,
					Message:     "Un-expected error occurs",
					Description: "An un-expected error has occurred",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetError(tt.args.errContext, tt.args.errMessage, tt.args.errCode)
			if got.Code != tt.want.Code || got.ErrorInfor.Code != tt.want.ErrorInfor.Code {
				t.Fatalf("Want %q, got %q", tt.want, got)
			}
		})
	}
}
