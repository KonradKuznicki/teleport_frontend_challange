import React from 'react';
import { ComponentMeta, ComponentStory } from '@storybook/react';
import { ThemeProvider } from 'styled-components';
import { theme } from '../general/Elements';
import { FileDetails } from './FileDetails';

export default {
    title: 'Files/FileDetails',
    component: FileDetails,
} as ComponentMeta<typeof FileDetails>;

const Template: ComponentStory<typeof FileDetails> = () => (
    <ThemeProvider theme={theme}>
        <FileDetails />
    </ThemeProvider>
);

export const Default = Template.bind({});

Default.args = {};
