package ldap

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-ldap/ldap/v3"
	go_ldap "github.com/go-ldap/ldap/v3"
)

const (
	REASON_PASSWORD_EXPIRED = "Password expired"
)

var (
	LDAP_URL   = "ldap://localhost:389"
	connection *go_ldap.Conn
)

func init() {
	// Initing here - to minimize changes to the whole project
	Init()
}
func Init() {
	v, ok := os.LookupEnv("LDAP_URL")
	if ok {
		LDAP_URL = v
	}

	l, err := go_ldap.DialURL(LDAP_URL)
	if err != nil {
		log.Printf("LDAP auth will fail: connect to LDAP server error: %v", err)
		return
	}

	connection = l

	// connection.Debug = true
}
func Close() {
	if connection != nil {
		connection.Close()
	}
	connection = nil
}

func IsUnauthorized(w http.ResponseWriter, r *http.Request) bool {
	return !IsAuthorized(w, r)
}

func IsAuthorized(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "", http.StatusUnauthorized)
		return false
	}

	if connection == nil {
		log.Printf("LDAP auth error: not connected to LDAP server")
		http.Error(w, "", http.StatusUnauthorized)
		return false
	}

	dn := "uid=" + username + ",ou=users,dc=example,dc=com"
	res, err := Bind(dn, password)
	if err != nil {
		logError(w, res, err)
		return false
	}

	return true
}

func Bind(dn string, password string) (*go_ldap.SimpleBindResult, error) {
	controls := []ldap.Control{}
	pr := ldap.NewControlBeheraPasswordPolicy()
	controls = append(controls, pr)

	req := ldap.NewSimpleBindRequest(dn, password, controls)

	return connection.SimpleBind(req)
}

func logError(w http.ResponseWriter, res *go_ldap.SimpleBindResult, err error) {
	reason := ""
	if isReasonPasswordExpired(res) {
		reason = REASON_PASSWORD_EXPIRED
	}

	log.Printf("LDAP auth error: %v (%v)", err, reason)
	http.Error(w, reason, http.StatusUnauthorized)
}

func isReasonPasswordExpired(res *go_ldap.SimpleBindResult) bool {
	if len(res.Controls) == 0 {
		return false
	}

	if len(res.Controls[0].String()) == 0 {
		return false
	}

	r := res.Controls[0].String()
	return strings.Contains(r, REASON_PASSWORD_EXPIRED)
}
