import { Given, Then, When, Before } from 'cypress-cucumber-preprocessor/steps';

const url = 'https://localhost:3001';

Before(() => {
    cy.viewport(1600, 1200);
    cy.visit(url + '/login');
    cy.get('#login').type('user1');
    cy.get('#pass').type('pass1');
    cy.contains('Log In').click();
});
Given('I open file manager', () => {
    // nop
});

Then(/^I see (.+) in the body$/, (txt) => {
    cy.get('body').should('include.text', txt);
});

Then(/^I see (\d+) files?$/, (filesCount) => {
    cy.get('tr').should('have.length', filesCount + 1);
});

When(/^I click (.+)$/, (clicked) => {
    cy.contains(clicked).click();
});

Then(/^I see than I am in (.+) folder$/, (folder) => {
    cy.location('pathname').should('equal', '/files/' + folder);
    cy.get('ul').contains(folder);
});
