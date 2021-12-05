import { Given, Then } from 'cypress-cucumber-preprocessor/steps';

const url = 'http://localhost:3000';
Given('I open file manager', () => {
    cy.visit(url);
    cy.viewport(1600, 1200);
});

Then(/^I see (.*) in the body$/, (txt) => {
    cy.get('body').should('include.text', txt);
});
