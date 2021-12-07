import React from 'react';
import styled from 'styled-components';
import { Cell, Row, Title } from '../general/Elements';
import { User } from '../Auth/User';
import { FileType } from './FileType';

export const PageHeadContainer = styled.div`
  margin: 30px auto;
  color: ${(props) => props.theme.colors.primary2};
  max-width: ${(props) => props.theme.contentMasWidth};
  vertical-align: middle;

`;


export function FileDetailsPageHead() {
    return <PageHeadContainer>
        <Row>
            <Cell>
                <Title style={{ textAlign: 'left', fontSize: '25px' }}><FileType type='txt' /> image.txt</Title>
            </Cell>
            <Cell style={{ textAlign: 'right', verticalAlign: 'bottom' }}>
                <User />
            </Cell>
        </Row>
    </PageHeadContainer>;
}