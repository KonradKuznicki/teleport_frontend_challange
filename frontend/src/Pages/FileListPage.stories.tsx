import React from 'react';
import { ComponentMeta, ComponentStory } from '@storybook/react';
import { ThemeProvider } from 'styled-components';
import { Page, theme } from '../general/Elements';
import { PageHead } from '../Files/PageHead';
import { FilesList } from '../Files/FileList';


function ListPage() {
    return <Page>
        <PageHead />
        <FilesList />
    </Page>;
}

export default {
    title: 'Pages/FilesListPage',
    component: ListPage,
} as ComponentMeta<typeof ListPage>;

const Template: ComponentStory<typeof ListPage> = () => <ThemeProvider theme={theme}><ListPage /></ThemeProvider>;

export const Default = Template.bind({});

Default.args = {};
