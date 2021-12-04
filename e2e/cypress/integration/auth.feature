Feature: Authentication

  @focus
  Scenario: not authenticated user accesses protected resource
    Given I open file manager
    And I am unauthenticated user
    Then I am redirected to login form
