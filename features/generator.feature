Feature: Generators return a value generating iterator
  Scenario: Generator returns a GeneratingIterator that calls the provided GeneratorFunc repeat count times.
    Given a GeneratorFunc that returns the count and repeat concatenated with a comma.
    And a repeat value of 3
    When Generate() is called
    Then calling Next() until false is returned should return the following strings:
    | 0,3 |
    | 1,3 |
    | 2,3 |

  Scenario Outline: Sequence generates correct sequence of values
    Given a start value of <start>
    And an end value of <end>
    When Sequence is called
    Then calling Next() until false is returned should return the following values: "<results>"

    Examples:
      | start | end | results        |
      | -2    | 0   | -2,-1,0        |
      | 2     | 0   | 2,1,0          |
      | 0     | 3   | 0,1,2,3        |
      | -2    | 2   | -2,-1,0,1,2    |
      | 1     | -3  | 1,0,-1,-2,-3   |

  Scenario Outline: StepSequence generates correct sequence of values
    Given a start value of <start>
    And an end value of <end>
    And an step value of <step>
    When StepSequence is called
    Then calling Next() until false is returned should return the following values: "<results>"

    Examples:
      | start | end | step | results        |
      | -8    | 8   | 3    | -8,-5,-2,1,4,7 |
      | -8    | 8   | -3   | -8,-5,-2,1,4,7 |
      | 0     | 8   | 3    | 0,3,6          |
      | 0     | 8   | -3   | 0,3,6          |
      | 0     | -8  | 3    | 0,-3,-6        |
      | 0     | -8  | -3   | 0,-3,-6        |
      | -4    | 4   | 2    | -4,-2,0,2,4    |
      | -4    | 4   | -2   | -4,-2,0,2,4    |
      | -2    | 0   | 1    | -2,-1,0        |
      | -2    | 0   | -1   | -2,-1,0        |
      | 2     | 0   | 1    | 2,1,0          |
      | 2     | 0   | -1   | 2,1,0          |
      | 0     | 3   | 1    | 0,1,2,3        |
      | 0     | 3   | -1   | 0,1,2,3        |
