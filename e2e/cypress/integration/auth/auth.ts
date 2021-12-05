import { Given, Then, When } from 'cypress-cucumber-preprocessor/steps';

const url = 'https://localhost:3001';
Given('I open Home page', () => {
    cy.visit(url);
});

Given(/^I am unauthenticated user$/, function () {
    // nop
});

Then(/^I am redirected to login form$/, function () {
    cy.location('pathname').should('equal', '/login');
});

Then(/^I am redirected to file manager$/, function () {
    cy.location('pathname').should('equal', '/files');
});

Given(/^I open file manager$/, function () {
    cy.visit(url + '/files');
});

///
Given(/^I open login form$/, function () {
    cy.visit(url + '/login');
});

Then(/^I see (.*) in the body$/, (txt) => {
    cy.get('body').should('include.text', txt);
});

///
When(/^I submit valid login form$/, function () {
    cy.contains('log me in').click();
});
Then(/^I am authenticated$/, function () {
    // @ts-ignore
    cy.get('body').not('include.text', 'files');
});

function LogMeIn() {
    cy.contains('log me in').click();
}

///
Given(/^I am authenticated user$/, function () {
    cy.visit(url);
    LogMeIn();
});
When(/^I log out$/, function () {
    cy.visit(url + '/user/logout');
});
Then(/^I am unauthenticated$/, function () {
    cy.visit(url + '/files');
    cy.location('pathname').should('equal', '/login');
});
