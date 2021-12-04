import { Given, Then } from "cypress-cucumber-preprocessor/steps";

const url = 'http://localhost:3000'
Given('I open Home page', () => {
    cy.visit(url)
})

Then(`I see {string} in the body`, (txt) => {
    cy.get('p').should('include.text', txt);
})
Then(/^I am redirected to file manager$/, function() {
    cy.location('pathname').should('be', '/files')
});