Feature: Super Admin bootstrapping
  As the system
  I Want to create the initial super Admin via environment variables
  So that the system has a super admin on startup

  Scenario: System fails to start if any SA env var is missing
    Given the environment variables are not Set or are not complete
      | SA_NAME     | System Admin  |
      | SA_PASSWORD | strOnP@ssword |
    When the system starts up
    Then the system should retun error message informing the var missing
    And system should not star

  Scenario: System creates initial SA on first startup
    Given the environment variables on startup
      | SA_NAME     | System Admin  |
      | SA_EMAIL    | sa@system.com |
      | SA_PASSWORD | strOnP@ssword |
    And no active super admin exists in the system
    When the system startup
    Then a super admin should be created with email "sa@system.com"
    And the super admin status should be "ACTIVE"

  Scenario: System skips SA creation if active SA already exists
    Given the environment variables on startup
      | SA_NAME     | System Admin  |
      | SA_EMAIL    | sa@system.com |
      | SA_PASSWORD | strOnP@ssword |
    And an active super admin exists in the system
    When the system starts up
    Then no new super admin should be created
    And the existing super admin should remain "ACTIVE"