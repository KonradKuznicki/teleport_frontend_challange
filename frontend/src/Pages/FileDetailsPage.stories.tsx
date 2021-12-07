import React from 'react';
import { ComponentMeta, ComponentStory } from '@storybook/react';
import { ThemeProvider } from 'styled-components';
import { Page, theme } from '../general/Elements';
import { FileDetails } from '../Files/FileDetails';
import { FileDetailsPageHead } from '../Files/FileDetailsPageHead';


function FilePage(props: any) {
    return <Page style={{ display: 'table-cell' }} {...props}>
        <FileDetailsPageHead />
        <FileDetails />
    </Page>;
}

export default {
    title: 'Pages/FileDetailsPage',
    component: FilePage,
} as ComponentMeta<typeof FilePage>;

const Template: ComponentStory<typeof FilePage> = (args) => <ThemeProvider
    theme={theme}><FilePage {...args} /></ThemeProvider>;

export const Default = Template.bind({});

Default.args = { isError: false };

