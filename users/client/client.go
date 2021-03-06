package client

import (
	"context"
	"github.com/farhadf/micromovies2/users/pb"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	"github.com/farhadf/micromovies2/users"
	"google.golang.org/grpc/metadata"
)

func NewGRPCClient(conn *grpc.ClientConn) users.Service {
	var newUserEndpoint = grpctransport.NewClient(
		conn, "pb.Users", "NewUser",
		users.EncodeGRPCNewUserRequest,
		users.DecodeGRPCNewUserResponse,
		pb.NewUserResponse{},
		//this will inject context fields specified in injectContext func in the metadata to be passed to downstream service
		grpctransport.ClientBefore(injectContext),
	).Endpoint()
	var getUserByEmailEndpoint = grpctransport.NewClient(
		conn, "pb.Users", "GetUserByEmail",
		users.EncodeGRPCGetUserByEmailRequest,
		users.DecodeGRPCGetUserByEmailResponse,
		pb.GetUserByEmailResponse{},
		//this will inject context fields specified in injectContext func in the metadata to be passed to downstream service
		grpctransport.ClientBefore(injectContext),
	).Endpoint()
	var changePasswordEndpoint = grpctransport.NewClient(
		conn, "pb.Users", "ChangePassword",
		users.EncodeGRPCChangePasswordRequest,
		users.DecodeGRPCChangePasswordResponse,
		pb.ChangePasswordResponse{},
		//this will inject context fields specified in injectContext func in the metadata to be passed to downstream service
		grpctransport.ClientBefore(injectContext),
	).Endpoint()
	var loginEndpoint = grpctransport.NewClient(
		conn, "pb.Users", "Login",
		users.EncodeGRPCLoginRequest,
		users.DecodeGRPCLoginResponse,
		pb.LoginResponse{},
		//this will inject context fields specified in injectContext func in the metadata to be passed to downstream service
		grpctransport.ClientBefore(injectContext),
	).Endpoint()
	return users.Endpoints{
		NewUserEndpoint:        newUserEndpoint,
		GetUserByEmailEndpoint: getUserByEmailEndpoint,
		ChangePasswordEndpoint: changePasswordEndpoint,
		LoginEndpoint:          loginEndpoint,
	}
}

func NewUser(ctx context.Context, service users.Service, user users.User) (string, error) {
	id, err := service.NewUser(ctx, user)
	if err != nil {
		return "", err
	}
	return id, nil
}

func GetUserByEmail(ctx context.Context, service users.Service, email string) (users.User, error) {
	user, err := service.GetUserByEmail(ctx, email)
	if err != nil {
		return users.User{}, err
	}
	return user, nil
}

func ChangePassword(ctx context.Context, service users.Service, email string, currentPassword string, newPassword string) (bool, error) {
	success, err := service.ChangePassword(ctx, email, currentPassword, newPassword)
	if err != nil {
		return success, err
	}
	return success, nil
}

func Login(ctx context.Context, service users.Service, email string, Password string) (string, string, error) {
	token, refreshToken, err := service.Login(ctx, email, Password)
	if err != nil {
		return "", "", err
	}
	return token, refreshToken, nil
}

//client before function to inject context into grpc metadata to pass to downstream service
func injectContext(ctx context.Context, md *metadata.MD) context.Context {
	if email, ok := ctx.Value("email").(string); ok {
		(*md)["email"] = append((*md)["email"], email)
	}
	if role, ok := ctx.Value("role").(string); ok {
		(*md)["role"] = append((*md)["role"], role)
	}
	if correlationid, ok := ctx.Value("correlationid").(string); ok {
		(*md)["correlationid"] = append((*md)["correlationid"], correlationid)
	}
	return ctx
}
