import React from 'react';
import styled, { keyframes } from 'styled-components';
import { CustomIcon } from '../general/Elements';


const UploadIcon = styled(CustomIcon).attrs({ code: 'upload' })`
  font-size: 50px;
`;

const UserContainer = styled.div`
  display: inline-block;
  font-family: ${(props) => props.theme.font};
`;

const fadeIn = keyframes`
  0% {
    opacity: 0;
  }
  100% {
    opacity: 1;
  }
`;

const FadeInButton = styled.span`
  font-family: ${(props) => props.theme.font};
  font-size: 25px;
  opacity: 0;
  margin: 0;
  cursor: pointer;

  ${UserContainer}:hover & {
    animation: 0.2s ${fadeIn} ease-out;
    opacity: 1;
  }

  &:hover {
    text-decoration: underline;
  }
`;


export function Upload() {
    return <UserContainer>
        <FadeInButton>Upload</FadeInButton>
        <UploadIcon />
    </UserContainer>;
}