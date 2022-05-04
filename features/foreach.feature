Feature: Foreach takes an Iterable and calls a function with each element

  Scenario: An Iterable with int 1,2, & 3 items is processed 3 times
    Given an Iterable with the following values:
      | 1 |
      | 2 |
      | 3 |
    And a foreach function that sums and counts the calls
    When Foreach is called
    Then The returned sum is 6
    Then The returned count is 3