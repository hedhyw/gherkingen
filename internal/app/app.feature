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
    | <feature>   | <template>                 |
    | app.feature | ../assets/std.args.v1.go.tmpl |
    | app.feature | @/std.args.v1.go.tmpl         |
    | app.feature | @/std.struct.v1.go.tmpl       |