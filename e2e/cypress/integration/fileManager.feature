Feature: File Manager

  Scenario: File manager loads
    When I open file manager
    Then I see notes.txt in the body

  Scenario: listing
    When I open file manager
    Then I see 4 files

  Scenario: go to foler
    Given I open file manager
    When I click images
    Then I see than I am in images folder
    And I see 2 files
