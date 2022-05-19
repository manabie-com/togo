package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	db "task-manage/internal/db/sqlc"
	"task-manage/internal/utils"
	"time"
)

type CreateUserParams struct {
	UserName         string `json:"user_name" binding:"required"`
	Password         string `json:"password" binding:"required"`
	MaximumTaskInDay int32  `json:"maximum_task_in_day" binding:"required"`
}

func (server *Server) createUser(context *gin.Context) {
	var req CreateUserParams
	var err error
	if err = context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	var hashPass string
	hashPass, err = utils.HashPassword(req.Password)
	arg := db.CreateUserParams{
		UserName:         req.UserName,
		HashedPassword:   hashPass,
		MaximumTaskInDay: req.MaximumTaskInDay,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
	var user db.User
	user, err = server.queries.CreateUser(context, arg)
	if err != nil {
		context.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	context.JSON(http.StatusOK, userResponse{
		Username:         user.UserName,
		CreatedAt:        user.CreatedAt,
		MaximumTaskInDay: user.MaximumTaskInDay,
	})
}

type listUserRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) getListUsers(ctx *gin.Context) {
	var req listUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	args := db.GetUsersParams{
		Offset: (req.PageID - 1) * req.PageSize,
		Limit:  req.PageSize,
	}
	users, err := server.queries.GetUsers(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var result []userResponse
	for _, user := range users {
		result = append(result, userResponse{
			Id:               user.ID,
			Username:         user.UserName,
			CreatedAt:        user.CreatedAt,
			MaximumTaskInDay: user.MaximumTaskInDay,
		})
	}
	ctx.JSON(http.StatusOK, result)
}

type userResponse struct {
	Id                int32     `json:"id"`
	Username          string    `json:"user_name"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
	MaximumTaskInDay  int32     `json:"maximum_task_in_day"`
}
type loginUserRequest struct {
	Username string `json:"user_name" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=3"`
}

type loginUserResponse struct {
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.queries.GetUserByName(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = utils.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		user.UserName,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.UserName,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := loginUserResponse{
		//SessionID:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
	}
	ctx.JSON(http.StatusOK, rsp)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
