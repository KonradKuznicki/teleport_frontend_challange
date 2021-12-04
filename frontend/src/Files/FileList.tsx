import React from 'react';
import styled from 'styled-components';
import { Box, Cell, CustomIcon, Row } from '../general/Elements';
import { Upload } from './Upload';
import { Path } from './Path';
import { FileType } from './FileType';

const ListBox = styled(Box)`
    width: 100%;
    max-width: ${(props) => props.theme.contentMasWidth};
    margin: auto;
`;

const TH = styled.th`
    font-family: ${(props) => props.theme.font};
    font-weight: ${(props) => props.theme.fontWeightBolder};
    text-align: left;
    cursor: pointer;
`;

const TD = styled.td``;

const TR = styled.tr`
    padding: 2px;
    border: 1px solid ${(props) => props.theme.colors.primary2};
    cursor: pointer;

    :hover {
        font-weight: 500;
    }
`;

const THead = styled.thead`
    tr {
        border-bottom: 1px solid ${(props) => props.theme.colors.darker1};
    }
`;

const lastUsed = new Date()
    .toISOString()
    .replace('T', ' ')
    .replace(/\..+$/, '');
const files = [
    { name: 'images', type: 'folder', size: '4 items', date: lastUsed },
    { name: 'mountains.jpg', type: 'jpg', size: '7MB', date: lastUsed },
    { name: 'test.pdf', type: 'PDF', size: '12KB', date: lastUsed },
    { name: 'some_file.pdf', type: 'PDF', size: '12KB', date: lastUsed },
    {
        name: 'stuff with spaces.pdf',
        type: 'PDF',
        size: '12KB',
        date: lastUsed,
    },
    { name: 'lol.pdf', type: 'PDF', size: '12KB', date: lastUsed },
];

export function FilesList() {
    return (
        <ListBox>
            <Row style={{ margin: '0' }}>
                <Cell>
                    <Path parts={['documents']} />
                </Cell>
                <Cell style={{ textAlign: 'right' }}>
                    <Upload />
                </Cell>
            </Row>

            <table style={{ width: '100%', borderCollapse: 'collapse' }}>
                <THead>
                    <tr>
                        <TH style={{ width: '40px' }} />
                        <TH>Name</TH>
                        <TH>Type</TH>
                        <TH>
                            Size <CustomIcon code="arrow_drop_down" />
                        </TH>
                    </tr>
                </THead>
                <tbody>
                    {files.map(({ name, type, size }) => (
                        <TR>
                            <TD>
                                <FileType type={type} />
                            </TD>
                            <TD>{name}</TD>
                            <TD>{type}</TD>
                            <TD>{size}</TD>
                        </TR>
                    ))}
                </tbody>
            </table>
        </ListBox>
    );
}
