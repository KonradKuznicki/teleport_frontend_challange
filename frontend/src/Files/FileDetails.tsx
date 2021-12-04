import React from 'react';
import styled from 'styled-components';
import { Box, Cell, Row } from '../general/Elements';
import { Path } from './Path';

const FileBox = styled(Box)`
  width: 100%;
  max-width: ${({theme}) => theme.contentMasWidth};
  margin: auto;
`;

function Property({ name, value }: { name: string, value: any }) {
    return <Row>
        <Cell style={{ fontWeight: 500, textTransform: 'capitalize' }}>{name}:</Cell>
        <Cell>{value}</Cell>
    </Row>;
}

const lastUsed = (new Date()).toISOString().replace('T', ' ').replace(/\..+$/, '');
const props = {
    name: 'test.pdf', type: 'PDF', size: '12KB', date: lastUsed,
};

export const FileDetails = () => <FileBox>
    <Path parts={['Documents', 'image.txt']} />
    {Object.entries(props).map(([k, v]) => <Property name={k} value={v} />)}
</FileBox>;