import React from 'react';
import { ComponentMeta, ComponentStory } from '@storybook/react';
import { ThemeProvider } from 'styled-components';
import { Page, theme } from '../general/Elements';
import { Alert } from '../general/Alert';
import { LoginForm } from '../Auth/LoginForm';


function LoginPage(props: { isError?: boolean }) {
    return <Page style={{ display: 'table-cell' }}>
        {props.isError && <Alert type='danger'>Incorrect Username and/or password</Alert>}
        <LoginForm />
    </Page>;
}

export default {
    title: 'Pages/LoginPage',
    component: LoginPage,
} as ComponentMeta<typeof LoginPage>;

const Template: ComponentStory<typeof LoginPage> = (args) => <ThemeProvider
    theme={theme}><LoginPage {...args} /></ThemeProvider>;

export const Default = Template.bind({});

Default.args = { isError: false };

export const Error = Template.bind({});

Error.args = { isError: true };
