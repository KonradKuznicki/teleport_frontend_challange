import React from 'react';
import styled from 'styled-components';
import { CustomIcon } from './Elements';


const Closer = styled(CustomIcon).attrs({ code: 'close' })`
  font-size: 50px;
  display: inline-block;
  vertical-align: middle;
  margin-right: 20px;
`;

const AlertContainer = styled.div`
  text-align: center;
  padding: 40px;
`;

const AlertContent = styled.div`
  font-family: ${(props) => props.theme.font};
  font-size: 30px;
  display: inline;
  vertical-align: middle;
`;

function AlertStructure(props: any) {

    return <AlertContainer {...props} >
        <Closer />
        <AlertContent>{props.children}</AlertContent>
    </AlertContainer>;
}

export const Alert = styled(AlertStructure)`
  color: ${props => props.theme.colors[props.type]};
  width: 100%;
  display: block;
  max-width: ${(props) => props.theme.contentMasWidth};
  margin: auto;
  cursor: pointer;

`;