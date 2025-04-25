package errors

import "testing"

func TestCustomError(t *testing.T) {
    tests := []struct {
        name       string
        code       int
        message    string
        wantString string
    }{
        {
            name:       "Not Found Error",
            code:       404,
            message:    "Resource not found",
            wantString: "Error 404: Resource not found",
        },
        {
            name:       "Unauthorized Error",
            code:       401,
            message:    "Unauthorized access",
            wantString: "Error 401: Unauthorized access",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := New(tt.code, tt.message)
            if err.Error() != tt.wantString {
                t.Errorf("CustomError.Error() = %v, want %v", err.Error(), tt.wantString)
            }
        })
    }
}

func TestIsNotFound(t *testing.T) {
    tests := []struct {
        name string
        err  error
        want bool
    }{
        {
            name: "Is Not Found",
            err:  New(404, "Not found"),
            want: true,
        },
        {
            name: "Is Not Not Found",
            err:  New(500, "Internal error"),
            want: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := IsNotFound(tt.err); got != tt.want {
                t.Errorf("IsNotFound() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestIsUnauthorized(t *testing.T) {
    tests := []struct {
        name string
        err  error
        want bool
    }{
        {
            name: "Is Unauthorized",
            err:  New(401, "Unauthorized"),
            want: true,
        },
        {
            name: "Is Not Unauthorized",
            err:  New(500, "Internal error"),
            want: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := IsUnauthorized(tt.err); got != tt.want {
                t.Errorf("IsUnauthorized() = %v, want %v", got, tt.want)
            }
        })
    }
}
