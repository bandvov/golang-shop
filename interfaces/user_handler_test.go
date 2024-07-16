package interfaces

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/bandvov/golang-shop/application"
	"github.com/bandvov/golang-shop/domain/users"
	"github.com/jackc/pgx"
)

func TestUserHandler_GetUsers(t *testing.T) {

	u := []*users.User{
		{ID: 1, FirstName: "John", LastName: "Doe", Status: "active", Role: "admin", Email: "john.doe@example.com"},
		{ID: 2, FirstName: "Jane", LastName: "Smith", Status: "active", Role: "admin", Email: "jane.smith@example.com"},
		{ID: 3, FirstName: "Alice", LastName: "Johnson", Status: "active", Role: "admin", Email: "alice.johnson@example.com"},
	}

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	tests := []struct {
		name       string
		h          *UserHandler
		args       args
		wantStatus int
		want       interface{}
	}{
		{
			name: "Get users successfully",
			h: NewUserHandler(application.NewUserService(&users.MockUserRepository{
				GetUsersFunc: func(ctx context.Context) ([]*users.User, error) {
					return u, nil
				},
			})),
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/users", nil),
			},
			wantStatus: http.StatusOK,
			want:       u, // Expecting []*users.User
		},
		{
			name: "Get users error",
			h: NewUserHandler(application.NewUserService(&users.MockUserRepository{
				GetUsersFunc: func(ctx context.Context) ([]*users.User, error) {
					return nil, pgx.ErrNoRows
				},
			})),
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/users", nil),
			},
			wantStatus: http.StatusInternalServerError,
			want:       pgx.ErrNoRows.Error(), // Expecting error message
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.GetUsers(tt.args.w, tt.args.r)

			resp := tt.args.w.(*httptest.ResponseRecorder)

			// Check status code
			if tt.wantStatus != resp.Code {
				t.Errorf("Status is not correct. Wanted: %+v, got: %+v", tt.wantStatus, resp.Code)
			}

			// Read response body
			body := resp.Body

			// Compare response body
			switch want := tt.want.(type) {
			case []*users.User:
				// Marshal expected []*users.User to JSON
				wantJSON, err := json.Marshal(want)
				if err != nil {
					t.Fatalf("Failed to marshal expected data: %v", err)
				}

				if bytes.Equal(body.Bytes(), wantJSON) {
					t.Errorf("Body is not correct. Wanted: %v. Got: %v", wantJSON, body)
				}
			case string:
				if strings.TrimSuffix(body.String(), "\n") != want {
					t.Errorf("Body is not correct. Wanted: %s. Got: %s", want, body)
				}
			default:
				t.Errorf("Unexpected type for tt.want: %T", tt.want)
			}
		})
	}
}

func TestUserHandler_GetUserByID(t *testing.T) {
	u := &users.User{ID: 1, FirstName: "John", LastName: "Doe", Status: "active", Role: "admin", Email: "john.doe@example.com", CreatedAt: time.Now()}
	fmt.Println(fmt.Sprintf("/users/%d", u.ID))
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *UserHandler
		args args
		want interface{}
	}{
		{
			name: "get user",
			h: NewUserHandler(application.NewUserService(&users.MockUserRepository{
				GetByIDFunc: func(ctx context.Context, id int) (*users.User, error) {
					return u, nil
				},
			})),
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/users/1", nil),
			},
			want: u, // Expecting *users.User
		},
		{
			name: "Get user error",
			h: NewUserHandler(application.NewUserService(&users.MockUserRepository{
				GetUsersFunc: func(ctx context.Context) ([]*users.User, error) {
					return nil, pgx.ErrNoRows
				},
			})),
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest("GET", "/users/1", nil),
			},
			want: pgx.ErrNoRows.Error(), // Expecting error message
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.GetUserByID(tt.args.w, tt.args.r)

			resp := tt.args.w.(*httptest.ResponseRecorder)

			// Compare response body
			switch want := tt.want.(type) {
			case *users.User:
				fmt.Println(resp.Body.Bytes())
				// Marshal expected []*users.User to JSON
				wantJSON, err := json.Marshal(want)
				if err != nil {
					t.Fatalf("Failed to marshal expected data: %v", err)
				}
				fmt.Println(resp.Body.Bytes())
				if bytes.Equal(resp.Body.Bytes(), wantJSON) {
					t.Errorf("Body is not correct. Wanted: %v. Got: %v", wantJSON, resp.Body.Bytes())
				}
			case string:
				if strings.TrimSuffix(resp.Body.String(), "\n") != want {
					t.Errorf("Body is not correct. Wanted: %s. Got: %s", want, resp.Body.String())
				}
			default:
				t.Errorf("Unexpected type for tt.want: %T", tt.want)
			}

		})
	}
}
