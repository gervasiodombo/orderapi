package user_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cucumber/godog"
	config "github.com/oderapi/configs"
	"github.com/oderapi/src/infra/persistence"
	"github.com/oderapi/src/main/factory"
	"github.com/oderapi/src/main/factory/db"
	"github.com/oderapi/src/main/factory/usecase"
	"github.com/oderapi/tests/mocks"
	"gorm.io/gorm"
)

func init() {
	os.Setenv("APP_ENV", "test")
	if err := config.LoadProfile(); err != nil {
		panic(err)
	}
}

var testDatabase *mocks.TestDatabase

func TestMain(m *testing.M) {
	var err error
	testDatabase, err = mocks.StartPostgres(nil)
	if err != nil {
		panic(err)
	}
	defer testDatabase.Cleanup()
	if err := persistence.RunMigrations(testDatabase.DB); err != nil {
		panic(fmt.Sprintf("failed to run migrations: %v", err))
	}
	os.Exit(m.Run())
}

type saBootstrapContext struct {
	db      *gorm.DB
	err     error
	message string
}

// ─── Given ───────────────────────────────────────────────────────────────────
func (s *saBootstrapContext) theFollowingEnvironmentVariablesAreSet(table *godog.Table) error {
	for _, row := range table.Rows {
		os.Setenv(row.Cells[0].Value, row.Cells[1].Value)
	}
	return nil
}

func (s *saBootstrapContext) theSANameEnvironmentVariableIsNotSet() error {
	os.Unsetenv("SA_NAME")
	return nil
}

func (s *saBootstrapContext) theSAEmailEnvironmentVariableIsNotSet() error {
	os.Unsetenv("SA_EMAIL")
	return nil
}

// ─── When ────────────────────────────────────────────────────────────────────

func (s *saBootstrapContext) theSystemStartsUp() error {
	input, err := usecase.MakeBootstrapSaInput()
	if err != nil {
		s.err = err
		return nil
	}

	postgresDb := db.MakePostgresDB()
	bootstrapSa := usecase.MakeBootstrapSa(postgresDb)
	s.message, s.err = factory.MakeRunBootstrapSa(input, bootstrapSa)
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
	s := &saBootstrapContext{
		db: testDatabase.DB,
	}
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		s.db.Exec(`DELETE FROM "T_USER_ROLES"`)
		s.db.Exec(`DELETE FROM "T_USERS"`)

		os.Unsetenv("SA_NAME")
		os.Unsetenv("SA_EMAIL")
		os.Unsetenv("SA_USERNAME")
		os.Unsetenv("SA_PASSWORD")

		s.err = nil
		s.message = ""
		return ctx, nil
	})

	// Given
	ctx.Step(`^the following environment variables are set:$`, s.theFollowingEnvironmentVariablesAreSet)
	ctx.Step(`^the SA_NAME environment variable is not set$`, s.theSANameEnvironmentVariableIsNotSet)
	ctx.Step(`^the SA_EMAIL environment variable is not set$`, s.theSAEmailEnvironmentVariableIsNotSet)

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
