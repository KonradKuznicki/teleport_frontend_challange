import { Given, Then, When } from "cypress-cucumber-preprocessor/steps";

const url = 'http://localhost:3000'
Given('I open Home page', () => {
    cy.visit(url)
});

Given(/^I am unauthenticated user$/, function() {
    // nop
});

Then(/^I am redirected to login form$/, function() {
    cy.location('pathname').should('equal', '/login')
});

Then(/^I am redirected to file manager$/, function() {
    cy.location('pathname').should('equal', '/files')
});

Given(/^I open file manager$/, function() {
    cy.visit(url + "/files")
});

///
Given(/^I open login form$/, function() {
    cy.visit(url + "/login");
});

Then(/^I see (.*) in the body$/, (txt) => {
    cy.get('body').should('include.text', txt);
});

///
When(/^I submit valid login form$/, function() {
    cy.contains('log me in').click()
});
Then(/^I am authenticated$/, function() {
    cy.get('body').not('include.text', "login");

});