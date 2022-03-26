Feature: Application command line tool
  Scenario: User wants to generate the output in given format
    When <format> is given
    And  <feature> is provided
    Then the output should be generated
    Examples:
    | <feature>        | <format> | <assertion> |
    | app.feature      | go       | does        |
    | app.feature      | json     | does        |
    | app.feature      | raw      | does        |
    | app.feature      | invalid  | does not    |
    | notfound.feature | raw      | does not    |

  Scenario: User wants to see usage information
    When <flag> is provided
    Then usage should be printed
    Examples:
    | <flag> |
    | --help |

  Scenario: User wants to list built-in templates
    When <flag> is provided
    Then templates should be printed
    Examples:
    | <flag> |
    | --list |

  Scenario: User wants to use custom template
    When <template> is provided
    And  <feature> is provided
    Then the output should be generated
    Examples:
    | <feature>   | <template>                      |
    | app.feature | ../assets/std.struct.v1.go.tmpl |
    | app.feature | @/std.struct.v1.go.tmpl         |
  
  Scenario: User wants to set custom package
    When <package> is provided
    Then the output should contain <package>
    Examples:
    | <package>     |
    | app_test      |
    | example_test  |
  
  Scenario: User wants to generate a permanent json output
    When -format is json
    And -permanent-ids is <TheSameIDs>
    Then calling generation twice will produce the same output <TheSameIDs>
    Examples:
    | <TheSameIDs> |
    | true         |
    | false        |

  Scenario: User gives an invalid flag
    When flag -invalid is provided
    Then a generation failed
