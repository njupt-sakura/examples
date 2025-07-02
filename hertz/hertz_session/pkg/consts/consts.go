package consts

const (
	TCP              = "tcp"
	UserTableName    = "users"
	RedisAddr        = "127.0.0.1:6379"
	MaxIdleNum       = 10
	RedisPasswd      = ""
	SessionSecretKey = "session-secret"
	CsrfSecretKey    = "csrf-secret"
	CsrfKeyLookUp    = "form:csrf"
	Username         = "username"
	HertzSession     = "hertz-session"
	MysqlDefaultDsn  = "whoami:pass@tcp(localhost:3307)/db0?charset=utf8&parseTime=True&loc=Local"
)

const (
	Success     = "success"
	RegisterErr = "user already exists"
	LoginErr    = "wrong username or password"
	PageErr     = "please login first"
	CsrfErr     = "csrf exception"
)
