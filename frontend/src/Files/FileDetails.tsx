import React, { ReactNode } from 'react';
import styled from 'styled-components';
import { Box, Cell, Row } from '../general/Elements';
import { Path } from './Path';
import { Size } from './Size';
import { FileStats } from './FileList';

const FileBox = styled(Box)`
    width: 100%;
    max-width: ${({ theme }) => theme.contentMasWidth};
    margin: auto;
`;

function Property({
    name,
    children,
    ...props
}: {
    name: string;
    children?: ReactNode;
}) {
    return (
        <Row {...props}>
            <Cell style={{ fontWeight: 500, textTransform: 'capitalize' }}>
                {name}:
            </Cell>
            <Cell>{children}</Cell>
        </Row>
    );
}

export function FileDetails({
    pathParts,
    details,
}: {
    pathParts: string[];
    details: FileStats;
}) {
    return (
        <FileBox>
            <Path parts={pathParts} />
            <Property name="Name">{details.name}</Property>
            <Property name="Type">{details.type}</Property>
            <Property name="Size">
                <Size size={details.size} folder={details.type === 'folder'} />
            </Property>
        </FileBox>
    );
}
