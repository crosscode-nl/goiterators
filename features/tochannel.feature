Feature: ToChannel sends values from an Iterable to a channel

  Scenario: An Iterable with int 1,2, & 3 items sends 1, 2 and 3 to a channel
    Given an Iterable with the following values:
      | 1 |
      | 2 |
      | 3 |
    And a channel
    When ToChannel is called
    Then the following values are received on the channel
      | 1 |
      | 2 |
      | 3 |
