import React from 'react';
import styled, { css } from 'styled-components';
import { Box } from '../general/Elements';
import { Path } from './Path';
import { FileType } from './FileType';
import { useNavigate } from 'react-router-dom';
import { Size } from './Size';

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

export type FileStats = {
    name: string;
    size: number;
    type: string;
};

const StyledSortable = styled.div`
    ${(args: { sorts: boolean }) =>
        args.sorts &&
        css`
            font-weight: 800;
        `}
`;

function Sortable(props: {
    sortBy: string;
    column: string;
    onSort: (sortBy: string) => void;
    children: React.ReactNode;
}) {
    return (
        <StyledSortable
            sorts={props.column == props.sortBy}
            onClick={() => props.onSort(props.column)}
        >
            {props.children}
        </StyledSortable>
    );
}

export function FilesList({
    files,
    path,
    onSort,
    sortBy,
}: {
    path: string[];
    files: FileStats[];
    onSort: (sortBy: string) => void;
    sortBy: string;
}) {
    const navigate = useNavigate();
    return (
        <ListBox>
            <Path parts={path} />

            <table style={{ width: '100%', borderCollapse: 'collapse' }}>
                <THead>
                    <tr>
                        <TH style={{ width: '40px' }} />
                        <TH>
                            <Sortable
                                column="name"
                                sortBy={sortBy}
                                onSort={onSort}
                            >
                                Name
                            </Sortable>
                        </TH>
                        <TH>
                            <Sortable
                                column="type"
                                sortBy={sortBy}
                                onSort={onSort}
                            >
                                Type
                            </Sortable>
                        </TH>
                        <TH>
                            <Sortable
                                column="size"
                                sortBy={sortBy}
                                onSort={onSort}
                            >
                                Size
                            </Sortable>
                        </TH>
                    </tr>
                </THead>
                <tbody>
                    {files.map(({ name, type, size }) => (
                        <TR
                            key={name}
                            onClick={() =>
                                navigate(
                                    '/files/' + path.concat([name]).join('/'),
                                )
                            }
                        >
                            <TD>
                                <FileType type={type} />
                            </TD>
                            <TD>{name}</TD>
                            <TD>{type}</TD>
                            <TD>
                                <Size size={size} folder={type === 'folder'} />
                            </TD>
                        </TR>
                    ))}
                </tbody>
            </table>
        </ListBox>
    );
}
