import React from 'react';
import { ComponentMeta, ComponentStory } from '@storybook/react';
import { ThemeProvider } from 'styled-components';
import { Page, theme } from './Elements';
import { Alert } from './Alert';


export default {
    title: 'general/Alerts',
    component: Alert,
} as ComponentMeta<typeof Alert>;

const Template: ComponentStory<typeof Alert> = (args) => <ThemeProvider theme={theme}><Page><Alert
    type={args.type}>{args.content}</Alert></Page></ThemeProvider>;

export const Default = Template.bind({});
export const Error = Template.bind({});
export const Waring = Template.bind({});
export const Success = Template.bind({});


const content = 'quick brown fox something something...';

Default.args = { content, type: 'info' };
Error.args = { content, type: 'danger' };
Waring.args = { content, type: 'careful' };
Success.args = { content, type: 'good' };
