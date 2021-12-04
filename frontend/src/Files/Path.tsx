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

const DefaultRoot: string[] = ['My Files'];

export function Path({ parts: parts, ...props }: { parts: string[] }) {

    const items = DefaultRoot
        .concat(parts);

    const parents = items.slice(0, -1);
    const current = items[items.length - 1];

    const pathPart = (item: string) => <><Item><A href={'/' + item}>{item}</A></Item><Item>/</Item></>;

    return <List {...props}>{parents.map(pathPart)}<Item>{current}</Item></List>;
}