import { Given, Then, When } from 'cypress-cucumber-preprocessor/steps';

const url = 'https://localhost:3001';
Given('I open Home page', () => {
    cy.visit(url);
});

Given(/^I am unauthenticated user$/, function () {
    // nop
});

Then(/^I am redirected to login form$/, function () {
    cy.location('pathname').should('include', '/login');
});

Then(/^I am redirected to file manager$/, function () {
    cy.location('pathname').should('include', '/files');
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
    cy.get('#login').type('user1');
    cy.get('#pass').type('pass1');
    cy.contains('Log In').click();
});
Then(/^I am authenticated$/, function () {
    cy.location('pathname').should('include', '/files');

    // @ts-ignore
    cy.get('body').not('include.text', 'files');
});

function LogMeIn() {
    cy.request('POST', url + '/API/v1/user/login', {
        login: 'user1',
        pass: 'pass1',
    });
    cy.visit(url);
}

///
Given(/^I am authenticated user$/, function () {
    cy.visit(url);
    LogMeIn();
});
When(/^I log out$/, function () {
    cy.visit(url + '/API/v1/user/logout');
});
Then(/^I am unauthenticated$/, function () {
    cy.visit(url + '/files');
    cy.location('pathname').should('include', '/login');
});
When(/^I wait (\d+) seconds$/, function (wait) {
    cy.wait(wait * 1000);
});
