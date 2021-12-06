import React from 'react';
import styled from 'styled-components';
import { A } from '../general/Elements';

const List = styled.ul`
    margin: 0;
    padding: 0;
`;

const Item = styled.li`
    display: inline-block;
    margin: 5px;
`;

type PathPart = {
    name: string;
    path: string;
};

const DefaultRoot: PathPart[] = [{ name: 'My Files', path: '/' }];

export function Path({ parts, ...props }: { parts: string[] }) {
    const items = DefaultRoot.concat(
        parts.map((i, c, a) => ({
            name: i,
            path: (c ? a[c - 1] : '') + '/' + i,
        })),
    );

    const parents = items.slice(0, -1);
    const current = items[items.length - 1];

    const pathPart = ({ name, path }: PathPart) => (
        <>
            <Item key={name + 'i'}>
                <A key={name + 'a'} to={'/files' + path}>
                    {name}
                </A>
            </Item>
            <Item key={name + '-slash'}>/</Item>
        </>
    );

    return (
        <List {...props} key="ul">
            {parents.map(pathPart)}
            <Item key="current">{current.name}</Item>
        </List>
    );
}
