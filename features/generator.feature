Feature: Generators return a value generating iterator
  Scenario: Generator returns a GeneratingIterator that calls the provided GeneratorFunc repeat count times.
    Given a GeneratorFunc that returns the count and repeat concatenated with a comma.
    And a repeat value of 3
    When Generate() is called
    Then calling Next() until false is returned should return the following strings:
    | 0,3 |
    | 1,3 |
    | 2,3 |

  Scenario: StepSequence generates an increasing sequence of values
    Given a start value of 0
    And an end value of 3
    And an step value of 1
    When StepSequence is called
    Then calling Next() until false is returned should return the following integers:
    | 0 |
    | 1 |
    | 2 |
    | 3 |

  Scenario: Sequence generates an increasing sequence of values
    Given a start value of 0
    And an end value of 3
    When Sequence is called
    Then calling Next() until false is returned should return the following integers:
      | 0 |
      | 1 |
      | 2 |
      | 3 |

  Scenario: StepSequence generates an increasing sequence of values
    Given a start value of 2
    And an end value of 0
    And an step value of 1
    When StepSequence is called
    Then calling Next() until false is returned should return the following integers:
      | 2 |
      | 1 |
      | 0 |

  Scenario: StepSequence generates an increasing sequence of values
    Given a start value of 2
    And an end value of 0
    And an step value of -1
    When StepSequence is called
    Then calling Next() until false is returned should return the following integers:
      | 2 |
      | 1 |
      | 0 |

  Scenario: Sequence generates an increasing sequence of values
    Given a start value of 2
    And an end value of 0
    When Sequence is called
    Then calling Next() until false is returned should return the following integers:
      | 2 |
      | 1 |
      | 0 |


  Scenario: StepSequence generates an increasing sequence of values
    Given a start value of -2
    And an end value of 0
    And an step value of -1
    When StepSequence is called
    Then calling Next() until false is returned should return the following integers:
      | -2 |
      | -1 |
      | 0 |

  Scenario: Sequence generates an increasing sequence of values
    Given a start value of -2
    And an end value of 0
    When Sequence is called
    Then calling Next() until false is returned should return the following integers:
      | -2 |
      | -1 |
      | 0 |

  Scenario: StepSequence generates an increasing sequence of values
    Given a start value of -4
    And an end value of 4
    And an step value of -2
    When StepSequence is called
    Then calling Next() until false is returned should return the following integers:
      | -4 |
      | -2 |
      | 0 |
      | 2 |
      | 4 |


  Scenario: StepSequence generates an increasing sequence of values
    Given a start value of 0
    And an end value of 8
    And an step value of 3
    When StepSequence is called
    Then calling Next() until false is returned should return the following integers:
      | 0 |
      | 3 |
      | 6 |

  Scenario: StepSequence generates an increasing sequence of values
    Given a start value of 0
    And an end value of -8
    And an step value of 3
    When StepSequence is called
    Then calling Next() until false is returned should return the following integers:
      | 0 |
      | -3 |
      | -6 |

  Scenario: StepSequence generates an increasing sequence of values
    Given a start value of -8
    And an end value of 8
    And an step value of 3
    When StepSequence is called
    Then calling Next() until false is returned should return the following integers:
      | -8 |
      | -5 |
      | -2 |
      | 1 |
      | 4 |
      | 7 |
