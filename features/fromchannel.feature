Feature: FromChannel returns an iterator
  A valid Iterator and functioning iterator is returned when FromChannel is called

  Scenario: A channel with 3 items send to it returns an Iterable that returns exactly 3 items when FromChannel is called
    Given a closed channel with the following values:
      | 1 |
      | 2 |
      | 3 |
    When FromChannel is called
    Then Next() returns true 3 times and then returns false

  Scenario: A channel with 3 items send to it returns an Iterable that returns the provided items when FromChannel is called
    Given a closed channel with the following values:
      | 1 |
      | 2 |
      | 3 |
    When FromChannel is called
    Then calling Next() until false is returned should return the following integers:
      | 1 |
      | 2 |
      | 3 |




