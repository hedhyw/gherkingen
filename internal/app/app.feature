Feature: Application command line tool

  Scenario Outline: User wants to generate the output in given format
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

  Scenario Outline: User wants to see usage information
    When <flag> is provided
    Then usage should be printed
    Examples:
      | <flag> |
      | --help |

  Scenario Outline: User wants to list built-in templates
    When <flag> is provided
    Then templates should be printed
    Examples:
      | <flag> |
      | --list |

  Scenario Outline: User wants to use custom template
    When <template> is provided
    And  <feature> is provided
    Then the output should be generated
    Examples:
      | <feature>   | <template>                      |
      | app.feature | ../assets/std.simple.v1.go.tmpl |
      | app.feature | @/std.simple.v1.go.tmpl         |

  Scenario Outline: User wants to set custom package
    When <package> is provided
    Then the output should contain <package>
    Examples:
      | <package>    |
      | app_test     |
      | example_test |

  Scenario Outline: User wants to generate a permanent json output
    When -format is json
    And -permanent-ids is <TheSameIDs>
    Then calling generation twice will produce the same output <TheSameIDs>
    Examples:
      | <TheSameIDs> |
      | true         |
      | false        |

  Scenario: User provides an invalid flag
    When flag -invalid is provided
    Then an error is returned

  Scenario Outline: User asks for a version
    When <flag> is provided
    Then version is printed
    Examples:
      | <flag>    |
      | --version |
      | -version  |

  Scenario Outline: User specifies a file, but the file is not found
    When inexistent <template> is provided
    And  <feature> is provided
    Then the user receives an error
    Examples:
      | <feature>   | <template> |
      | app.feature | not_found  |

  Scenario: User wants to run tests in parallel
    When `scenario.feature` is given
    Then generated code contains `t.Parallel()`

  Scenario: User wants to run tests sequentially
    When `-disable-go-parallel` is provided
    And `scenario.feature` is given
    Then generated code doesn't contain `t.Parallel()`

  Scenario Outline: User wants to generate the output for a feature written in a specific natural language
    When the <language> is given
    And the <feature> is provided
    Then the output should be generated
    Examples:
      | <language>  | <feature>                                      | assertion |
      | en          | ../generator/examples/simple.feature           | does      |
      | en-pirate   | ../generator/examples/simple.en-pirate.feature | does      |
      | unsupported | app.feature                                    | does not  |

  Scenario Outline: User wants to see all supported natural languages
    When the <flag> is provided
    Then the list of supported natural languages should be printed
    Examples:
      | <flag>      |
      | -languages  |
      | --languages |

  Scenario: User wants to see consistent supported natural languages output
    When the user lists supported natural languages several times
    Then the output is the same

  Scenario Outline: User wants to be able to specify the natural language of a feature in its file name
    When the <language> is specified by the flag
    And the <feature> is given
    Then the output should be generated
    Examples:
      | <language> | <feature>                                  | assertion |
      | en-pirate  | ./testdata/pirate.feature                  | does      |
      |            | ./testdata/pirate.feature                  | does not  |
      |            | ./testdata/pirate.sample.en-pirate.feature | does      |
      | wrong      | ./testdata/pirate.sample.en-pirate.feature | does not  |
      | en         | ./testdata/english.feature                 | does      |
      |            | ./testdata/english.feature                 | does      |
      |            | ./testdata/english.sample.en.feature       | does      |
      | wrong      | ./testdata/english.sample.en.feature       | does not  |
      |            | app.feature                                | does      |
