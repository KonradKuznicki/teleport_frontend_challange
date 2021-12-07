Feature: Home page

  Home page does not exist yet so redirect to file manager

  @focus
  Scenario: redirecting home to file manager
    Given I open Home page
    Then I am redirected to file manager
