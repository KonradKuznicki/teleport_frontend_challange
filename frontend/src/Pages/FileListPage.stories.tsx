import React from 'react';
import { ComponentMeta, ComponentStory } from '@storybook/react';
import { ThemeProvider } from 'styled-components';
import { theme } from '../general/Elements';
import { ListFilesPage } from './ListFilesPage';
import { FileStats } from '../Files/FileList';

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

export default {
    title: 'Pages/FilesListPage',
    component: ListFilesPage,
} as ComponentMeta<typeof ListFilesPage>;

const Template: ComponentStory<typeof ListFilesPage> = (args) => (
    <ThemeProvider theme={theme}>
        <ListFilesPage path={args.path} files={args.files} />
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
