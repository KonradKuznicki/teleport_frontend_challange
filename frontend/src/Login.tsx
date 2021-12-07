import React, { useEffect, useState } from 'react';
import ReactDOM from 'react-dom';
import { createGlobalStyle, ThemeProvider } from 'styled-components';
import { Page, theme, Title } from './general/Elements';
import { LoginForm } from './Auth/LoginForm';
import { Alert } from './general/Alert';
import './index.css';

createGlobalStyle`
  @import url('https://fonts.googleapis.com/css2?family=M+PLUS+2:wght@800&display=swap');

  html, body {
    height: 100%;
    min-height: 100%;

    margin: 0;
    padding: 0;

    font-family: ${theme.font};
  }

`;

async function login(payload: any) {
    return fetch('/API/v1/user/login', {
        body: JSON.stringify(payload),
        headers: {
            'Content-Type': 'application/json',
        },
        // credentials: 'init',
        method: 'POST',
    });
}

export function useFetchWithSubmit(): {
    data: any;
    isLoading: boolean;
    isError: false | string;
    submit: (payload: any) => void;
} {
    const [payload, setPayload] = useState(undefined as any);
    const [data, setData] = useState(undefined as any);
    const [isLoading, setIsLoading] = useState(true);
    const [isError, setIsError] = useState<false | string>(false);

    useEffect(() => {
        const fetchData = async () => {
            if (payload === undefined) {
                setIsError(false);
                setIsLoading(false);
                setData(undefined);
                return;
            }

            setIsError(false);
            setIsLoading(true);

            try {
                const result = await login(payload);
                const response = await result.json();

                if (response && response.goTo) {
                    console.log(response);
                    window.location.pathname = response.goTo;
                }
            } catch (error) {
                setIsError((error as Object).toString());
                setData({ error: 'something went wrong' });
            }

            setIsLoading(false);
        };

        fetchData();
    }, [payload]);

    return { data, isLoading, isError, submit: setPayload };
}

export function LoginPage() {
    const { isError, isLoading, data, submit } = useFetchWithSubmit();
    return (
        <Page style={{ display: 'table-cell' }}>
            {isError && (
                <Alert type="danger">
                    {isError} <br /> {data && data.error}
                </Alert>
            )}
            {!isLoading && (
                <LoginForm submit={(login, pass) => submit({ login, pass })} />
            )}
            {isLoading && <Title>loading...</Title>}
        </Page>
    );
}

ReactDOM.render(
    <React.StrictMode>
        <ThemeProvider theme={theme}>
            <LoginPage />
        </ThemeProvider>
    </React.StrictMode>,
    document.getElementById('root'),
);
