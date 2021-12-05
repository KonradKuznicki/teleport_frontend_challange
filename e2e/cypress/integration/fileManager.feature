Feature: File Manager

  Scenario: File manager loads
    When I open file manager
    Then I see mountains.jpg in the body
