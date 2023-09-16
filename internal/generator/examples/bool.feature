Feature: Type determinatiopn
  Scenario Outline: All type are determinated
    When generator completed
    Then correct types are shown
    Examples:
    | <bool> | <int> | <string> | <flag> | <float64> |
    | true   | 1     | hello    | -      | 1.0       |
    | false  | 2     | world    | +      | 0.0       |
