Feature: Interpret Hello World

    Scenario: Send Interpret Hello World with success 
    The user is going to develop
    When The user initialize the input with data `>++++++++[<+++++++++>-]<.>++++[<+++++++>-]<+.+++++++..+++.>>++++++[<+++++++>-]<++.------------.>++++++[<+++++++++>-]<+.<.+++.------.--------.>>>++++[<++++++++>-]<+` to the interpreter
    Then The output should be "Hello World!
