import React from 'react';
import styled, { keyframes } from 'styled-components';
import { CustomIcon } from '../general/Elements';


const AccountIcon = styled(CustomIcon).attrs({ code: 'person_outline' })`
    font-size: 30px;
`;

const UserContainer = styled.div`
  display: inline-block;
  color: ${(props) => props.theme.colors.primary2};
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
  color: ${(props) => props.theme.colors.primary2};
  font-family: ${(props) => props.theme.font};
  opacity: 0;
  // margin: 10px;
  cursor: pointer;

  ${UserContainer}:hover & {
    animation: 0.2s ${fadeIn} ease-out;
    opacity: 1;
  }

  &:hover {
    text-decoration: underline;
  }
`;


export function User() {
    return <UserContainer>
        <FadeInButton>Log Out</FadeInButton>
        <AccountIcon />
    </UserContainer>;
}