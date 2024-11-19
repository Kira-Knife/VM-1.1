package test

import (
	"context"
	"os"
	"testing"

	. "VictoriaMetrics/test/pkg/testgodoglib"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"

	. "github.com/onsi/gomega"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^Сервер OpenLDAP$`, серверOpenLDAP)
	ctx.Step(`^Политика паролей$`, политикаПаролей)
	ctx.Step(`^На сервере существует пользователь: "([^"]*)", с паролем "([^"]*)"$`, наСервереСуществуетПользовательСПаролем)
	ctx.Step(`^На сервере пользователь с истёкшим паролем: "([^"]*)"$`, наСервереПользовательСИстёкшимПаролем)
	ctx.Step(`^Другой адрес LDAP сервера: "([^"]*)"$`, другойАдресLDAPСервера)
	ctx.Step(`^Модуль авторизации подключается к LDAP серверу$`, модульАвторизацииПодключаетсяКLDAPСерверу)
	ctx.Step(`^Вызывается точка входа "([^"]*)": "([^"]*)", "([^"]*)"$`, вызываетсяТочкаВхода)
	ctx.Step(`^Получен успешный ответ$`, полученУспешныйОтвет)
	ctx.Step(`^Получена ошибка авторизации: "([^"]*)"$`, полученаОшибкаАвторизации)
	ctx.Step(`^Пауза$`, пауза)

	// -----------------------
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		beforeScenario()
		return ctx, nil
	})
	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		afterScenario()
		return ctx, nil
	})
	InitializeGomegaForGodog(ctx)
	_ = Ω
}

func InitializeSuite(tsc *godog.TestSuiteContext) {
	tsc.BeforeSuite(beforeSuite)
	tsc.AfterSuite(afterSuite)
}

func TestMain(m *testing.M) {
	var opts = godog.Options{
		Output:        colors.Colored(os.Stdout),
		Strict:        true,
		StopOnFailure: true,
	}

	godog.BindCommandLineFlags("godog.", &opts)
	r := godog.TestSuite{
		Name:                 "app",
		TestSuiteInitializer: InitializeSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()
	os.Exit(r)
}
