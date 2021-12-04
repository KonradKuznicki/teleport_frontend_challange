import React from 'react';
import { ComponentMeta, ComponentStory } from '@storybook/react';
import { ThemeProvider } from 'styled-components';
import { theme } from '../general/Elements';
import { FilesList } from './FileList';

export default {
    title: 'Files/FilesList',
    component: FilesList,
} as ComponentMeta<typeof FilesList>;

const Template: ComponentStory<typeof FilesList> = () => <ThemeProvider theme={theme}><FilesList /></ThemeProvider>;

export const Default = Template.bind({});

Default.args = {};
