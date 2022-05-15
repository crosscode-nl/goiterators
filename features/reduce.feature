Feature: Reduce takes an Iterable and a reduce function to reduce a collection to a single result

  Scenario: An Iterable with int 1,2, & 3 items is reduced so a sum of 6
    Given an Iterable with the following values:
      | 1 |
      | 2 |
      | 3 |
    And a reduce function that sums all values
    And initial value of 0
    When Reduce is called
    Then The returned sum is 6
