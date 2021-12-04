import React from 'react';
import { ComponentMeta, ComponentStory } from '@storybook/react';
import { ThemeProvider } from 'styled-components';
import { theme } from '../general/Elements';
import { User } from './User';

export default {
    title: 'Auth/UserWidget',
    component: User,
} as ComponentMeta<typeof User>;

const Template: ComponentStory<typeof User> = () => <ThemeProvider theme={theme}><User /></ThemeProvider>;

export const Default = Template.bind({});

Default.args = {};
