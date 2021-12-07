import React from 'react';
import { ComponentMeta, ComponentStory } from '@storybook/react';
import { ThemeProvider } from 'styled-components';
import { theme } from '../general/Elements';
import { FilesList, FileStats } from './FileList';

export default {
    title: 'Files/FilesList',
    component: FilesList,
} as ComponentMeta<typeof FilesList>;

const files: FileStats[] = [
    { name: 'images', type: 'folder', size: 4 },
    {
        name: 'mountains.jpg',
        type: 'jpg',
        size: 7 * Math.pow(1024, 2),
    },
    { name: 'test.pdf', type: 'PDF', size: 12000 },
    { name: 'some_file.pdf', type: 'PDF', size: 12000 },
    {
        name: 'stuff with spaces.pdf',
        type: 'PDF',
        size: 12000,
    },
    { name: 'lol.pdf', type: 'PDF', size: 12000 },
];

const Template: ComponentStory<typeof FilesList> = (args) => (
    <ThemeProvider theme={theme}>
        <FilesList path={args.path} files={args.files} />
    </ThemeProvider>
);

export const Default = Template.bind({});

Default.args = {
    path: [],
    files: [],
};

export const One = Template.bind({});

One.args = {
    path: ['documents'],
    files,
};
