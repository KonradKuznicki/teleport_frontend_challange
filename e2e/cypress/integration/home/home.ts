import { Given, Then } from "cypress-cucumber-preprocessor/steps";

const url = 'http://localhost:3000'

const visited = [];

Given('I open Home page', () => {
    cy.intercept('*', (t) => {
        visited.push(t.url);
    })
    cy.visit(url)

});

Then(/^I am redirected to file manager$/, function() {
    expect(visited[1]).to.eq(`${url}/files`);
});
