package apigateway

import (
	"context"
	"encoding/json"
	"net/http"
)

//decode function for each method
func decodeLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req loginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

//decode function for each method
func decodeRegisterRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req registerRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

//decode function for each method
func decodeChangePasswordRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req changePasswordRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

//decode function for each method
func decodeNewMovieRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req newMovieRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
