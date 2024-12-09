package testlib

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/cucumber/godog"
	. "github.com/onsi/gomega"
)

var Be = Equal

func NoErr(err error) {
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
}

var Ok = NoErr

func Err(err error) {
	ExpectWithOffset(1, err).To(HaveOccurred())
}
func ErrIs(err, target error) {
	ExpectWithOffset(1, err).To(HaveOccurred())
	v := errors.Is(err, target)
	s := fmt.Sprintf("ErrIs: '%v' != '%v'", err, target)
	ExpectWithOffset(1, v).To(BeTrue(), s)
}

func Yes(b bool) {
	ExpectWithOffset(1, b).To(BeTrue())
}
func YesText(b bool, text string) {
	ExpectWithOffset(1, b).To(BeTrue(), text)
}

func SetTestDeadlockProtection() {
	go checkTestTimeout(1 * time.Second)
}
func checkTestTimeout(t time.Duration) {
	<-time.After(t)
	panic("Test Suite Timeout")
}

func Atoi(in string) int {
	res, err := strconv.Atoi(in)
	Ok(err)
	return res
}
func Atof(in string) float64 {
	res, err := strconv.ParseFloat(in, 64)
	Ok(err)
	return res
}

var TestResult error

func InitializeGomegaForGodog(ctx *godog.ScenarioContext) {
	ctx.StepContext().Before(func(ctx context.Context, st *godog.Step) (context.Context, error) {
		TestResult = nil
		return ctx, nil
	})
	ctx.StepContext().After(func(ctx context.Context, st *godog.Step, status godog.StepResultStatus, err error) (context.Context, error) {
		return ctx, TestResult
	})
	RegisterFailHandler(func(message string, callerSkip ...int) {
		// remember only the expectation failed first
		// anything thereafter is not to be believed
		if TestResult == nil {
			TestResult = fmt.Errorf(message)
		}
	})
}
