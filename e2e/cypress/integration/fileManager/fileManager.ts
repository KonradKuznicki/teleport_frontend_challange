import { Given, Then, When } from 'cypress-cucumber-preprocessor/steps';

const url = 'http://localhost:3000';
Given('I open file manager', () => {
    cy.visit(url + '/files');
    cy.viewport(1600, 1200);
});

Then(/^I see (.+) in the body$/, (txt) => {
    cy.get('body').should('include.text', txt);
});

Then(/^I see (\d+) files?$/, (filesCount) => {
    cy.get('tr').should('have.length', filesCount);
});

When(/^I click (.+)$/, (clicked) => {
    cy.contains(clicked).click();
});

Then(/^I see than I am in (.+) folder$/, (folder) => {
    cy.location('pathname').should('equal', '/files/' + folder);
    cy.get('ul').contains(folder);
});
