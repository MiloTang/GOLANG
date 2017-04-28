package config

import (
	"webnet/comlib/session"
)

const Host = "127.0.0.1:3306"
const Username = "root"
const Passwd = ""
const Database = "test"
const Charset = "utf8"

var CS *session.CookieSession

func init() {
	session.CookieName = "beenoob"
	session.MaxLT = 3600
	session.Salt = "beenoob"
	CS = session.NewCookieSession()
}
