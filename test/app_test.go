package test

import (
	. "VictoriaMetrics/test/pkg/testgodoglib"
	. "VictoriaMetrics/test/pkg/testlib"
	"flag"
	"net/http"
	"net/http/httptest"
	"os"

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

func наСервереСуществуетПользовательСПаролем(user, pass string) {
	ДобавитьOpenLDAPПользователя()
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

func полученаОшибкаАвторизации() {
	Ω(t.resp.StatusCode).To(Be(401), "полученаОшибкаАвторизации")

	d := make([]byte, 10)
	n, err := t.resp.Body.Read(d)
	Ok(err)

	Ω(n).To(BeNumerically("<=", 1), "полученаОшибкаАвторизации: body length")
	_ = t.resp.Body.Close()
}

func полученУспешныйОтвет() {
	Ω(t.resp.StatusCode).To(BeElementOf([]int{200, 204}), "полученУспешныйОтвет")
	_ = t.resp.Body.Close()
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
        \
        -u 0 \
        \
        ` + OpenLDAPImageName + ` >/dev/null`)

	ЖдатьОткрытияАдресПорта(OpenLDAP_TCPAddr, 3000)

	if DEBUG_SHOW_DOCKER_LOGS {
		Bash(`sudo docker logs -f ` + OpenLDAPContainerName + ` &`)
	}
}

func ДобавитьOpenLDAPПользователя() {
	Bash(`sudo docker exec -i \
        ` + OpenLDAPContainerName + ` \
		/conf/add_new_user \
        `)
}

func ОстановитьOpenLDAP() {
	// Bash(`sudo docker logs ` + OpenLDAPContainerName)
	Bash(`sudo docker rm -f ` + OpenLDAPContainerName + "&>/dev/null")
}
