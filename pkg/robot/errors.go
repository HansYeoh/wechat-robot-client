package robot

import "errors"

var (
	// ErrForbidden 禁止当前账号登录
	ErrForbidden = errors.New("login forbidden")

	// ErrInvalidStorage define invalid storage error
	ErrInvalidStorage = errors.New("invalid storage")

	// NetworkErr define wechat network error
	ErrNetwork = errors.New("wechat network error")

	// ErrNoSuchUserFound define no such user found error
	ErrNoSuchUserFound = errors.New("no such user found")

	// ErrLoginTimeout define login timeout error
	ErrLoginTimeout = errors.New("login timeout")

	// ErrWebWxDataTicketNotFound define webwx_data_ticket not found error
	ErrWebWxDataTicketNotFound = errors.New("webwx_data_ticket not found")

	// ErrUserLogout define user logout error
	ErrUserLogout = errors.New("user logout")

	// ErrUserNotLogin define user not login
	ErrUserNotLogin = errors.New("user not login")
)

// Error impl error interface
func (r Ret) Error() string {
	return r.String()
}
