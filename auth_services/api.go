package auth_services

import (
	"context"
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"mini_project/common"
	"mini_project/config"
	db_api "mini_project/db"
	"mini_project/db/model"

	"sync"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/projects"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/tokens"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// implements the authServer with backend is keystone
type authServer struct {
	db model.DatabaseAPI
	//m      *sync.RWMutex
	client gophercloud.ServiceClient
	// default project
	defaultProject string
	m              *sync.RWMutex
}

var authPolicy embed.FS

func NewAuthServer(dbUrl map[string]string) (*authServer, error) {
	// connect mysql
	db, err := db_api.GetDatabase(dbUrl)

	if err != nil {
		return nil, err
	}
	client := gophercloud.ServiceClient{
		ProviderClient: &gophercloud.ProviderClient{},
		Endpoint:       config.AuthUrl(),
	}
	client.UseTokenLock()
	return &authServer{db: db, client: client, m: &sync.RWMutex{}}, err
}

func (s *authServer) Login(ctx context.Context, _ *empty.Empty) (*AuthResponse, error) {
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<< Login >>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	s.m.RLock()
	defer s.m.RUnlock()
	pair := ctx.Value("token").([]string)
	projectName := "admin"
	scope := tokens.Scope{
		ProjectName: projectName,
		DomainID:    "default",
	}

	authOpts := tokens.AuthOptions{
		Username: pair[0],
		Password: pair[1],
		DomainID: "default",
	}

	if pair[0] == projectName {
		authOpts.Scope = scope
	}
	// safety for concurrence
	//s.client.UseTokenLock()
	result := tokens.Create(&s.client, &authOpts)
	if result.Err != nil {
		return nil, status.Error(codes.Unauthenticated, result.Err.Error())
	}

	userInfo, err := result.ExtractUser()
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	token, err := result.ExtractToken()
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if s.defaultProject == "" && pair[0] == projectName {
		project, err := result.ExtractProject()
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
		s.defaultProject = project.ID
	}
	// timeNow := time.Now()

	var audit struct {
		AuditIDs []string `json:"audit_ids"`
	}
	err = result.ExtractInto(&audit)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	fields := map[string]interface{}{
		"name":    userInfo.Name,
		"enabled": true,
	}

	_, err = s.db.AddUser(userInfo.ID, fields)
	if err != nil {
		panic(err)
	}
	expiresAt, err := ptypes.TimestampProto(token.ExpiresAt)
	if err != nil {
		panic(err)
	}
	user, _ := s.db.GetUser(userInfo.ID)
	grpc.SendHeader(ctx, metadata.New(map[string]string{"X-Subject-Token": token.ID}))
	return &AuthResponse{
		TokenExpiresAt: expiresAt,
		Id:             userInfo.ID,
		Name:           userInfo.Name,
		UserName:       user.UserName,
		Enabled:        true,
	}, nil
	// return nil, nil
}

func (s *authServer) CreateUser(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<< CreateUser >>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	s.m.Lock()
	defer s.m.Unlock()
	detailErr := ValidateCreateUserRequest(req, s.db)
	if detailErr != nil {
		st := status.Newf(codes.InvalidArgument, "Invalid argument")
		st, _ = st.WithDetails(&DetailError{
			Old:    nil,
			New:    common.Interface2String(req),
			Errors: detailErr,
		})
		return nil, st.Err()
	}
	adminProject, err := s.getDefaultProjectID()
	if err != nil {
		panic(err)
	}

	createOpts := users.CreateOpts{
		Name:             req.Name,
		DomainID:         "default",
		DefaultProjectID: adminProject,
		Enabled:          &req.Enabled,
		Password:         req.Password,
		Extra: map[string]interface{}{
			"email":     req.Email,
			"phone":     req.Phone,
			"user_name": req.UserName,
		},
	}

	userInfo, err := users.Create(&s.client, createOpts).Extract()
	if err != nil {
		panic(err)
	}

	user := model.User{
		ID:       userInfo.ID,
		Name:     userInfo.Name,
		UserName: req.UserName,
		Email:    req.Email,
		Phone:    req.Phone,
		Enabled:  userInfo.Enabled,
	}
	err = s.db.CreateUser(user)
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	var resp UserResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		panic(err)
	}

	return &resp, nil
}

func (s *authServer) UpdateUser(ctx context.Context, req *UserUpdateReq) (*UserResponse, error) {
	s.m.Lock()
	defer s.m.Unlock()

	adminProject, err := s.getDefaultProjectID()
	if err != nil {
		panic(err)
	}
	current, err := s.db.GetUser(req.UserId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	currentUser := UserResponse{
		Id:       current.ID,
		Name:     current.Name,
		UserName: current.UserName,
		Email:    current.Email,
		Phone:    current.Phone,
		Enabled:  current.Enabled,
	}
	old := common.Interface2String(&currentUser)
	to := common.Interface2String(req)

	var updateOpts users.UpdateOpts
	fields := make(map[string]interface{})
	update := make(map[string]interface{})

	updateOpts = users.UpdateOpts{
		DomainID:         "default",
		DefaultProjectID: adminProject,
		Extra:            update,
	}

	userInfo, err := users.Update(&s.client, req.UserId, updateOpts).Extract()
	if err != nil {
		panic(err)
	}

	fields["user_name"] = fmt.Sprintf("%v", userInfo.Extra["user_name"])
	fields["email"] = fmt.Sprintf("%v", userInfo.Extra["email"])
	fields["phone"] = fmt.Sprintf("%v", userInfo.Extra["phone"])

	updated, _ := s.db.GetUser(req.UserId)
	var from map[string]string
	for k := range to {
		from[k] = old[k]
	}
	return &UserResponse{
		Id:       updated.ID,
		Name:     updated.Name,
		UserName: updated.UserName,
		Email:    updated.Email,
		Phone:    updated.Phone,
		Enabled:  updated.Enabled,
	}, nil
}

func (s *authServer) ResetPassword(ctx context.Context, req *ResetPasswdReq) (*empty.Empty, error) {
	s.m.Lock()
	defer s.m.Unlock()
	_, err := s.db.GetUser(req.UserId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	updateOpts := users.UpdateOpts{
		DomainID: "default",
		Password: req.Password,
	}
	_, err = users.Update(&s.client, req.UserId, updateOpts).Extract()
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *authServer) DeleteUser(ctx context.Context, req *UserDeleteReq) (*empty.Empty, error) {
	s.m.Lock()
	defer s.m.Unlock()
	_, err := s.db.GetUser(req.UserId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	adminProject, err := s.getDefaultProjectID()
	if err != nil {
		panic(err)
	}
	disable := false
	// disable user
	updateOpts := users.UpdateOpts{
		DomainID:         "default",
		DefaultProjectID: adminProject,
		Enabled:          &disable,
	}
	_, err = users.Update(&s.client, req.UserId, updateOpts).Extract()
	if err != nil {
		panic(err)
	}
	fields := map[string]interface{}{"deleted": true}
	_, err = s.db.UpdateUser(req.UserId, fields)
	if err != nil {
		panic(err)
	}
	return &empty.Empty{}, nil
}

func (s *authServer) DeleteUserPermanently(ctx context.Context, req *UserDeleteReq) (*empty.Empty, error) {
	s.m.Lock()
	defer s.m.Unlock()
	_, err := s.db.GetUser(req.UserId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}
	err = users.Delete(&s.client, req.UserId).ExtractErr()
	if err != nil {
		return nil, err
	}
	user := model.User{ID: req.UserId}
	err = s.db.DeleteUser(user)
	if err != nil {
		panic(err)
	}
	return &empty.Empty{}, nil
}

func (s *authServer) ChangePassword(ctx context.Context, req *UserChangePasswdReq) (*empty.Empty, error) {
	s.m.Lock()
	defer s.m.Unlock()

	_, err := s.db.GetUser(req.UserId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "user not found")
	}

	changePasswordOpts := users.ChangePasswordOpts{
		OriginalPassword: req.OriginalPassword,
		Password:         req.Password,
	}

	err = users.ChangePassword(&s.client, req.UserId, changePasswordOpts).ExtractErr()
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (s *authServer) getDefaultProjectID() (string, error) {
	if s.defaultProject != "" {
		return s.defaultProject, nil
	}
	listOpts := projects.ListOpts{
		Enabled: gophercloud.Enabled,
	}

	allPages, err := projects.List(&s.client, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allProjects, err := projects.ExtractProjects(allPages)
	if err != nil {
		panic(err)
	}

	for _, project := range allProjects {
		if project.Name == "admin" {
			s.defaultProject = project.ID
			return s.defaultProject, nil
		}
	}
	return "", errors.New("unknown")
}

func (s *authServer) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return common.VerifyTokenBearerOrBasic(ctx, fullMethodName, s.db, &s.client)
	// return nil, nil
}

// Close () Cleanup handle database
func (s *authServer) Close() error {
	if err := s.db.Close(); err != nil {
		return err
	}
	return nil
}
