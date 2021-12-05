import React from 'react';
import { ComponentMeta, ComponentStory } from '@storybook/react';
import { ThemeProvider } from 'styled-components';
import { theme } from '../general/Elements';
import { ListFilesPage } from './ListFilesPage';

export default {
    title: 'Pages/FilesListPage',
    component: ListFilesPage,
} as ComponentMeta<typeof ListFilesPage>;

const Template: ComponentStory<typeof ListFilesPage> = () => (
    <ThemeProvider theme={theme}>
        <ListFilesPage />
    </ThemeProvider>
);

export const Default = Template.bind({});

Default.args = {};
