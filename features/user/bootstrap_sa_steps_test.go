package user_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/oderapi/src/main/factory/usecase"
)

type saBootstrapContext struct {
	err error
}

// ─── Given ───────────────────────────────────────────────────────────────────
func (s *saBootstrapContext) theFollowingEnvironmentVariablesAreSet(table *godog.Table) error {
	for _, row := range table.Rows[1:] {
		os.Setenv(row.Cells[0].Value, row.Cells[1].Value)
	}
	return nil
}

func (s *saBootstrapContext) theSANameEnvironmentVariableIsNotSet() error {
	os.Unsetenv("SA_NAME")
	return nil
}

// ─── When ────────────────────────────────────────────────────────────────────

func (s *saBootstrapContext) theSystemStartsUp() error {
	_, s.err = usecase.MakeBootstrapSaInput()
	return nil
}

// ─── Then ────────────────────────────────────────────────────────────────────

func (s *saBootstrapContext) theSystemShouldReturnErrorMessage(message string) error {
	if s.err == nil {
		return fmt.Errorf("expected error, got nil")
	}
	if s.err.Error() != message {
		return fmt.Errorf("expected error message '%s', got '%s'", message, s.err.Error())
	}
	return nil
}

func (s *saBootstrapContext) theSystemShouldNotStart() error {
	if s.err == nil {
		return fmt.Errorf("expected system to not start, but no error was returned")
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	s := &saBootstrapContext{}
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		os.Setenv("SA_NAME", "System Admin")
		os.Setenv("SA_EMAIL", "sa@system.com")
		os.Setenv("SA_USERNAME", "super_admin")
		os.Setenv("SA_PASSWORD", "strOnP@ssword")
		return ctx, nil
	})

	// Given
	ctx.Step(`^the following environment variables are set:$`, s.theFollowingEnvironmentVariablesAreSet)
	ctx.Step(`^the SA_NAME environment variable is not set$`, s.theSANameEnvironmentVariableIsNotSet)

	// When
	ctx.Step(`^the system starts up$`, s.theSystemStartsUp)

	// Then
	ctx.Step(`^the system should return error message "([^"]*)"$`, s.theSystemShouldReturnErrorMessage)
	ctx.Step(`^the system should not start$`, s.theSystemShouldNotStart)
}

// ─── Ponto de entrada ─────────────────────────────────────────────────────────
func TestSABootstrapFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"bootstrap_sa.feature"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("SA Bootstrap BDD scenarios failed")
	}
}
