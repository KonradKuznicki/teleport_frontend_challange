import { Given, Then } from "cypress-cucumber-preprocessor/steps";

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
Given(/^I open file manager$/, function() {
    cy.visit(url + "/files")
});