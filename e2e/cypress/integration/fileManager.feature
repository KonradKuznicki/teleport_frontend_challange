Feature: File Manager

  Scenario: File manager loads
    When I open file manager
    Then I see mountains.jpg in the body

  Scenario: listing
    When I open file manager
    Then I see 7 files

  Scenario: go to foler
    Given I open file manager
    When I click images
    Then I see than I am in images folder
    And I see 1 file
