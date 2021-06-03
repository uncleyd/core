package server

type AuthCookie struct {
	Token       string
	AccountId   int64
	AccountName string
	Name        string
}

func (a *AuthCookie) Parse(ctx *GinContext) {
	a.Token, _ = ctx.Cookie("token")
	a.AccountId = ctx.CookieInt64("accountId", 0)
	a.AccountName, _ = ctx.Cookie("accountName")
	a.Name, _ = ctx.Cookie("name")
}

func (a *AuthCookie) IsValidate() bool {
	if a.Token == "" || a.AccountId < 1 || a.AccountName == "" || a.Name == "" {
		return false
	}

	return false
}
