import React from 'react';
import { ComponentMeta, ComponentStory } from '@storybook/react';
import { ThemeProvider } from 'styled-components';
import { theme } from '../general/Elements';
import { Path } from './Path';


export default {
    title: 'Files/Path',
    component: Path,
} as ComponentMeta<typeof Path>;

const Template: ComponentStory<typeof Path> = (args) => <ThemeProvider theme={theme}><Path
    parts={args.parts} /></ThemeProvider>;

export const Empty = Template.bind({});

Empty.args = {
    parts: [],
};

export const One = Template.bind({});

One.args = {
    parts: ['Documents'],
};


export const Many = Template.bind({});

Many.args = {
    parts: ['Documents', 'images', 'last summer', 'photo.jpg'],
};

