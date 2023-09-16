Feature: Guess the word
  Scenario Outline: <main_player> starts a game
    When the <main_player> starts a game at start_date
    Then the <main_player> waits for a <oponent_player> to join
    Examples:
    | <main_player> | <oponent_player> |
    | Mark          | Alex             |
    | Ivan          | Duong            |
