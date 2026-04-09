Feature: Super Admin Bootstrapping
  As the system
  I want to create the initial Super Admin via environment variables
  So that the system has a super admin on startup

  Background:
    Given the following environment variables are set:
      | SA_NAME     | System Admin  |
      | SA_EMAIL    | sa@system.com |
      | SA_USERNAME | super_admin   |
      | SA_PASSWORD | strOnP@ssword |

  Scenario: System returns error if SA_NAME env var is not set
    And the SA_NAME environment variable is not set
    When the system starts up
    Then the system should return error message "SA_NAME environment variable not set"
    And the system should not start

  Scenario: System returns error if SA_NAME env var is not set
    And the SA_NAME environment variable is not set
    When the system starts up
    Then the system should return error message "SA_NAME environment variable not set"
    And the system should not start