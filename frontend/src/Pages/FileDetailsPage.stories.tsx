import React from 'react';
import { ComponentMeta, ComponentStory } from '@storybook/react';
import { ThemeProvider } from 'styled-components';
import { theme } from '../general/Elements';
import { FileDetailsPage } from './FileDetailsPage';

export default {
    title: 'Pages/FileDetailsPage',
    component: FileDetailsPage,
} as ComponentMeta<typeof FileDetailsPage>;

const Template: ComponentStory<typeof FileDetailsPage> = (args) => (
    <ThemeProvider theme={theme}>
        <FileDetailsPage {...args} />
    </ThemeProvider>
);

export const Default = Template.bind({});

Default.args = {
    pathParts: ['docs'],
    details: { name: 'notes.txt', size: 123, type: 'pdf' },
};
