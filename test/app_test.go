package test

import (
	. "VictoriaMetrics/test/pkg/testgodoglib"
	. "VictoriaMetrics/test/pkg/testlib"
	"flag"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"time"

	"github.com/VictoriaMetrics/VictoriaMetrics/lib/auth/ldap"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/httpserver"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/ingestserver/opentsdbhttp"
	. "github.com/VictoriaMetrics/VictoriaMetrics/lib/lib"
	. "github.com/onsi/gomega"
)

const (
	PREFIX                 = "vm_test_"
	OpenLDAPImageName      = "my/openldap"
	OpenLDAPContainerName  = PREFIX + "openldap"
	SLAPD_PASSWORD         = "admin_pass"
	SLAPD_DOMAIN           = "example.com"
	OpenLDAP_TCPAddr       = "localhost:389"
	TestConfigsOpenLDAP    = "configs"
	DEBUG_SHOW_DOCKER_LOGS = false
)

type testData struct {
	resp *http.Response
}

var (
	t *testData
)

func resetTestData() {
	t = &testData{}
}

func beforeSuite() {
	flag.Parse()
}
func afterSuite() {
}

func beforeScenario() {
	resetTestData()
}
func afterScenario() {
	ОстановитьOpenLDAP()
	ldap.Close()
}

// -----------------------
func серверOpenLDAP() {
	ЗапуститьOpenLDAP()
}

func политикаПаролей() {
	OpenLDAPЗадатьПолитикуПаролей()
}

func наСервереСуществуетПользовательСПаролем(user, pass string) {
	ДобавитьOpenLDAPПользователя()
}
func наСервереПользовательСИстёкшимПаролем(arg1 string) {
	OpenLDAPПользовательсИстёкшимПаролем()
}

func другойАдресLDAPСервера(addr string) {
	os.Setenv("LDAP_URL", addr)
}

func модульАвторизацииПодключаетсяКLDAPСерверу() {
	ldap.Init()
}

func вызываетсяТочкаВхода(entry, user, pass string) {
	w := httptest.NewRecorder()

	switch entry {
	case "httpserver":
		req := httptest.NewRequest(http.MethodGet, "/-/ready", nil)
		req.SetBasicAuth(user, pass)
		httpserver.HandlerWrapper(nil, w, req, nil)
	case "ingestserver/opentsdbhttp":
		rh := opentsdbhttp.NewRequestHandler(func(r *http.Request) error {
			return nil
		})
		req := httptest.NewRequest(http.MethodPut, "/put", nil)
		req.SetBasicAuth(user, pass)
		rh.ServeHTTP(w, req)
	default:
		panic("entry?")
	}

	t.resp = w.Result()
}

func полученУспешныйОтвет() {
	Ω(t.resp.StatusCode).To(BeElementOf([]int{200, 204}), "полученУспешныйОтвет")
	_ = t.resp.Body.Close()
}

func полученаОшибкаАвторизации(msg string) {
	Ω(t.resp.StatusCode).To(Be(401), "полученаОшибкаАвторизации")

	d, err := io.ReadAll(t.resp.Body)
	Ok(err)
	defer t.resp.Body.Close()

	s := string(d)
	s = strings.TrimSpace(s)

	Ω(s).To(Be(msg), "полученаОшибкаАвторизации: msg")
}

// helpers
func ЗапуститьOpenLDAP() {
	BashV(`sudo docker run -d \
        --name=` + OpenLDAPContainerName + ` \
        --net=host \
        -v /etc/localtime:/etc/localtime:ro \
        \
        -v "${PWD}/` + TestConfigsOpenLDAP + `:/conf:ro" \
        \
        --tmpfs=/etc/ldap:rw,size=1000k,mode=777 \
        --tmpfs=/var/lib/ldap:rw,size=1000k,mode=777 \
        \
        -e SLAPD_PASSWORD=` + SLAPD_PASSWORD + ` \
        -e SLAPD_DOMAIN=` + SLAPD_DOMAIN + ` \
		-e SLAPD_ADDITIONAL_MODULES=ppolicy \
		\
        -u 0 \
		\
        ` + OpenLDAPImageName + ` >/dev/null`)

	if DEBUG_SHOW_DOCKER_LOGS {
		Bash(`sudo docker logs -f ` + OpenLDAPContainerName + ` &`)
	}

	ЖдатьОткрытияАдресПорта(OpenLDAP_TCPAddr, 3000)
}

func OpenLDAPЗадатьПолитикуПаролей() {
	Bash(`sudo docker exec -i \
		` + OpenLDAPContainerName + ` \
		/conf/set_password_policy \
		&> /dev/null \
		`)
}

func ДобавитьOpenLDAPПользователя() {
	Bash(`sudo docker exec -i \
        ` + OpenLDAPContainerName + ` \
		/conf/add_new_user \
        `)
}

func OpenLDAPПользовательсИстёкшимПаролем() {
	Bash(`sudo docker exec -i \
        ` + OpenLDAPContainerName + ` \
		/conf/expire_user_pass \
        `)
}

func ОстановитьOpenLDAP() {
	// Bash(`sudo docker logs ` + OpenLDAPContainerName)
	Bash(`sudo docker rm -f ` + OpenLDAPContainerName + "&>/dev/null")
}

func пауза() {
	time.Sleep(2000 * time.Millisecond)
}
