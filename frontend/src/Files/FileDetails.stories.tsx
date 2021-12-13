import React from 'react';
import { ComponentMeta, ComponentStory } from '@storybook/react';
import { ThemeProvider } from 'styled-components';
import { theme } from '../general/Elements';
import { FileDetails } from './FileDetails';

export default {
    title: 'Files/FileDetails',
    component: FileDetails,
} as ComponentMeta<typeof FileDetails>;

const Template: ComponentStory<typeof FileDetails> = (args) => (
    <ThemeProvider theme={theme}>
        <FileDetails {...args} />
    </ThemeProvider>
);

export const Default = Template.bind({});

Default.args = {
    pathParts: ['docs', 'images'],
    details: {
        size: 110330,
        type: 'pdf',
        name: 'some dokument.pdf',
    },
};
