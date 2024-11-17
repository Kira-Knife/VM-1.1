package ldap

import (
	"log"
	"net/http"
	"os"

	go_ldap "github.com/go-ldap/ldap/v3"
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
	err := connection.Bind(dn, password)
	if err != nil {
		log.Printf("LDAP auth error: %v", err)
		http.Error(w, "", http.StatusUnauthorized)
		return false
	}

	return true
}
