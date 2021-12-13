import React, { InputHTMLAttributes, useState } from 'react';
import styled, { keyframes } from 'styled-components';
import { Cell, CustomIcon, Input, Row } from '../general/Elements';
import { User } from '../Auth/User';

export const PageHeadContainer = styled.div`
    margin: 30px auto;
    color: ${(props) => props.theme.colors.primary2};
    max-width: ${(props) => props.theme.contentMasWidth};
    vertical-align: middle;
`;

const SearchIcon = styled(CustomIcon).attrs({ code: 'search' })`
    font-size: 30px;
    vertical-align: bottom;
`;

const SearchInput = styled(Input)`
    background: transparent;
    border: none;
    color: ${(props) => props.theme.colors.primary2};
    padding: 0;
    width: auto;
    font-size: 25px;

    &:focus {
        background: ${(props) => props.theme.colors.lighter1};
    }
`;

function SearchFN(props: InputHTMLAttributes<any>) {
    return (
        <div className={props.className}>
            <SearchIcon />
            <SearchInput
                id="search"
                placeholder={props.placeholder}
                onChange={props.onChange}
                value={props.value}
            />
        </div>
    );
}

const fadeIn = keyframes`
  0% {
    border-bottom-color: #181A23;
  }
  100% {
    border-bottom-color: #FFD764;
  }
`;

const Search = styled(SearchFN)`
    border-bottom: 2px solid #181a23;

    &:hover {
        animation: 0.2s ${fadeIn} ease-out;
        border-bottom: 2px solid #ffd764;
        opacity: 1;
    }

    ${SearchIcon} {
        margin-right: 10px;
    }
`;

export function PageHead({ onSearch }: { onSearch: (search: string) => void }) {
    const [state, setState] = useState('');
    onSearch(state);
    return (
        <PageHeadContainer>
            <Row>
                <Cell>
                    <Search
                        placeholder="<file name>"
                        onChange={(e) => setState(e.target.value)}
                        value={state}
                    />
                </Cell>
                <Cell style={{ textAlign: 'right', verticalAlign: 'bottom' }}>
                    <User />
                </Cell>
            </Row>
        </PageHeadContainer>
    );
}
