package ldap

import (
	"log"
	"net/http"
	"os"

	go_ldap "github.com/go-ldap/ldap/v3"
)

var (
	LDAP_URL = "ldap://localhost:389"
)

func init() {
	v, ok := os.LookupEnv("LDAP_URL")
	if ok {
		LDAP_URL = v
	}
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

	_ = go_ldap.DialURL
	l, err := go_ldap.DialURL(LDAP_URL)
	if err != nil {
		http.Error(w, "", http.StatusUnauthorized)
		return false
	}
	defer l.Close()

	dn := "uid=" + username + ",ou=users,dc=example,dc=com"
	err = l.Bind(dn, password)
	if err != nil {
		log.Printf("LDAP auth error: %v", err)
		http.Error(w, "", http.StatusUnauthorized)
		return false
	}

	return true
}
