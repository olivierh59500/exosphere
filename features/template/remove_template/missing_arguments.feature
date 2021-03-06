Feature: missing arguments

  Scenario: attempting to run 'exo template remove' without specifying template name
    Given I am in the root directory of an empty application called "test app"
    And my application is a Git repository
    When starting "exo template remove" in my application directory
    Then I eventually see:
      """
      wrong number of arguments
      """
    And it exits with code 1
