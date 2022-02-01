Feature: Interpret Hello World

    Scenario: Send Interpret Hello World with success 
    The user is going to develop
    When The user initialize the input with data "data/helloWord.bf" to the interpreter
    When The user compile the code
    Then The output should be "Hello, World"

    Scenario: test length instruction 
    The user is going to develop
    When The user initialize the input with data "data/helloWord.bf" to the interpreter
    When The user activate the operator  "+"
    Then The number of instruction should be 67

    Scenario: test Desactivate Operator 
    The user is going to develop
    When The user initialize the input with data "data/helloWord.bf" to the interpreter
    When The user desactivate the operator  "+"
    Then The number of instruction should be 50

    Scenario: test length instruction 
    The user is going to develop
    When The user initialize the input with data "data/helloWord.bf" to the interpreter
    When The user activate the operator  "+"
    Then The number of instruction should be 67
