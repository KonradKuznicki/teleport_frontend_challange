import React, { useState } from 'react';
import styled from 'styled-components';
import { Box, Button, Cell, Input, Label, Row } from '../general/Elements';

const LoginFormBox = styled(Box)`
    width: 340px;
    margin: auto;
`;

export function LoginForm({
    submit,
}: {
    submit: (login: string, pass: string) => void;
}) {
    const [login, setLogin] = useState('');
    const [pass, setPass] = useState('');
    return (
        <LoginFormBox>
            <Label>
                <Row>
                    <Cell>Login:</Cell>
                    <Cell>
                        <Input
                            value={login}
                            onChange={(e) => setLogin(e.target.value)}
                        />
                    </Cell>
                </Row>
            </Label>

            <Label>
                <Row>
                    <Cell>Password:</Cell>
                    <Cell>
                        <Input
                            type="password"
                            value={pass}
                            onChange={(e) => setPass(e.target.value)}
                        />
                    </Cell>
                </Row>
            </Label>

            <Row>
                <Cell> </Cell>
                <Cell style={{ textAlign: 'right' }}>
                    <Button onClick={() => submit(login, pass)}>Log In</Button>
                </Cell>
            </Row>
        </LoginFormBox>
    );
}
