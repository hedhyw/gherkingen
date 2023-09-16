Feature: Application command line tool
  Scenario Outline: User wants to see usage information
    When the application is started with <flag>
    Then usage should be printed <printed>
    And exit status should be <exit_status>
    Examples:
    | <flag>   | <exit_status> | <printed> |
    | --help   |       0       | true      |
    | -help    |       0       | true      |
    | -invalid |       1       | false     |
