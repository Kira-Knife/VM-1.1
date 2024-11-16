package test

import (
	. "VictoriaMetrics/test/pkg/testgodoglib"
	. "VictoriaMetrics/test/pkg/testlib"
	"flag"
	"net/http"
	"net/http/httptest"

	"github.com/VictoriaMetrics/VictoriaMetrics/lib/httpserver"
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
	// os.Setenv("LDAP_URL", "ldap://localhost:389")
}
func afterSuite() {
}

func beforeScenario() {
	resetTestData()
}
func afterScenario() {
	ОстановитьOpenLDAP()
}

// -----------------------
func серверOpenLDAP() {
	ЗапуститьOpenLDAP()
}

func наСервереСуществуетПользовательСПаролем(user, pass string) {
	ДобавитьOpenLDAPПользователя()
}

func вызываетсяТочкаВходаТребующаяАвторизации(user, pass string) {
	req := httptest.NewRequest(http.MethodGet, "/-/ready", nil)
	req.SetBasicAuth(user, pass)

	w := httptest.NewRecorder()
	httpserver.HandlerWrapper(nil, w, req, nil)

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
	Ω(t.resp.StatusCode).To(Be(200), "полученУспешныйОтвет")
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
