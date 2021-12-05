Feature: Authentication

  Scenario: not authenticated user accesses protected resource
    Given I open file manager
    And I am unauthenticated user
    Then I am redirected to login form

  Scenario: log in form
    Given I am unauthenticated user
    And I open login form
    Then I see login form in the body

  Scenario: log in process
    Given I am unauthenticated user
    And I open login form
    When I submit valid login form
    Then I am authenticated
    And I am redirected to file manager

  Scenario: log out
    Given I am authenticated user
    When I log out
    Then I am unauthenticated
    And I am redirected to login form


  Scenario: in out in
    Given I am authenticated user
    When I log out
    Then I am unauthenticated
    And I am redirected to login form
    When I submit valid login form
    Then I am authenticated
    And I am redirected to file manager
