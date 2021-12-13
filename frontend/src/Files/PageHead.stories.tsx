import React from 'react';
import { ComponentMeta, ComponentStory } from '@storybook/react';
import { ThemeProvider } from 'styled-components';
import { Page, theme } from '../general/Elements';
import { PageHead } from './PageHead';

function TestPage() {
    return (
        <Page>
            <PageHead onSearch={() => void 0} />
        </Page>
    );
}

export default {
    title: 'Files/PageHead',
    component: TestPage,
} as ComponentMeta<typeof TestPage>;

const Template: ComponentStory<typeof TestPage> = () => (
    <ThemeProvider theme={theme}>
        <TestPage />
    </ThemeProvider>
);

export const Default = Template.bind({});

Default.args = {};
