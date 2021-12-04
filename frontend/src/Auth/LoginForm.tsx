import React from 'react';
import styled from 'styled-components';
import { Box, Button, Cell, Input, Label, Row } from '../general/Elements';

const LoginFormBox = styled(Box)`
  width: 340px;
  margin: auto;
`;

export function LoginForm() {
    return <LoginFormBox>

        <Label><Row><Cell>Login:</Cell><Cell><Input /></Cell></Row></Label>

        <Label><Row><Cell>Password:</Cell><Cell><Input type='password' /></Cell></Row></Label>

        <Row>
            <Cell> </Cell>
            <Cell style={{ textAlign: 'right' }}>
                <Button>Log In</Button>
            </Cell>
        </Row>
    </LoginFormBox>;
}
