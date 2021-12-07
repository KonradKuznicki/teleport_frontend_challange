import React from 'react';
import { ComponentMeta, ComponentStory } from '@storybook/react';
import { ThemeProvider } from 'styled-components';
import { theme } from '../general/Elements';

import { LoginForm } from './LoginForm';

export default {
    title: 'Auth/LoginForm',
    component: LoginForm,
} as ComponentMeta<typeof LoginForm>;

const Template: ComponentStory<typeof LoginForm> = () => (
    <ThemeProvider theme={theme}>
        <LoginForm submit={() => void 0} />
    </ThemeProvider>
);

export const Default = Template.bind({});

Default.args = {};
