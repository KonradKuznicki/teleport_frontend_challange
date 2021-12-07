import React from 'react';
import { render, RenderResult } from '@testing-library/react';
import { Path } from './Path';
import { ThemeProvider } from 'styled-components';
import { theme } from '../general/Elements';

describe('Path', function () {
    describe('empty path', function () {
        const root = 'My Files';

        let node: RenderResult<
            typeof import('@testing-library/dom/types/queries'),
            HTMLElement
        >;
        beforeEach(() => {
            node = render(
                <ThemeProvider theme={theme}>
                    <Path parts={[]} />
                </ThemeProvider>,
            );
        });
        test('shows root', function () {
            expect(node.getByText(root)).toBeInTheDocument();
        });
        test('no slash when only root', function () {
            expect(node.queryAllByText('/')).toHaveLength(0);
        });
        it('should match snapshot', function () {
            expect(node.asFragment().firstChild).toMatchSnapshot();
        });
    });

    describe('non empty path', function () {
        const root = 'My Files';

        let node: RenderResult<
            typeof import('@testing-library/dom/types/queries'),
            HTMLElement
        >;
        beforeEach(() => {
            node = render(
                <ThemeProvider theme={theme}>
                    <Path parts={['documents', 'images']} />
                </ThemeProvider>,
            );
        });
        test('shows root', function () {
            expect(node.getByText(root)).toBeInTheDocument();
        });
        test('shows slashes between items', function () {
            expect(node.queryAllByText('/')).toHaveLength(2);
        });
        it('should match snapshot', function () {
            expect(node.asFragment().firstChild).toMatchSnapshot();
        });
    });
});
