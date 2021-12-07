import React from 'react';
import styled, { createGlobalStyle } from 'styled-components';
import { Link } from 'react-router-dom';

export const theme = {
    font: "'M PLUS 2', sans-serif",
    fontWeight: 400,
    fontWeightBolder: 500,
    contentMasWidth: '1000px',
    colors: {
        lighter1: '#181A23',
        primary1: '#08170C',
        darker1: '#000000',
        primary2: '#FFD764',
        lighter2: '#ffe7bc',
        lightest2: '#ffffff',
        primary3: '#afffb8',
        darker3: '#82fe8f',
        primary4: '#005de0',
        darker4: '#0010be',
        danger: '#ff001f',
        careful: '#ff6a00',
        good: '#57ff63',
        info: '#00ffff',
    },
};

createGlobalStyle`
  @import url('https://fonts.googleapis.com/css2?family=M+PLUS+2:wght@800&display=swap');
  @import url('https://fonts.googleapis.com/icon?family=Material+Icons'); // TODO: why it doesn't work?

  html, body {
    height: 100%;
    min-height: 100%;

    margin: 0;
    padding: 0;

    font-family: ${theme.font};
  }

`;

// font-family: 'Nunito Sans', 'Helvetica Neue', Helvetica, Arial, sans-serif;

export const Title = styled.h1`
    font-size: 30px;
    text-align: center;
    font-family: ${(props) => props.theme.font};
    font-weight: ${(props) => props.theme.fontWeightBolder};

    margin: 0;
    padding: 0;
`;

export const Box = styled.div`
    box-sizing: border-box;
    padding: 20px;
    background: ${(props) => props.theme.colors.primary2};
`;

export const Label = styled.label`
    font-family: ${(props) => props.theme.font};
    font-weight: ${(props) => props.theme.defaultFontWeight};
`;

export const Input = styled.input`
    border: none;
    border-bottom: 2px solid ${(props) => props.theme.colors.darker1};
    padding: 5px;
    font-size: 17px;
    box-sizing: border-box;
    background: ${(props) => props.theme.colors.lighter2};

    &:focus {
        outline: none;
        background: ${(props) => props.theme.colors.lightest2};
    }
`;

export const A = styled(Link)`
    color: ${(props) => props.theme.colors.primary4};
    font-family: ${(props) => props.theme.font};
    font-weight: ${(props) => props.theme.defaultFontWeight};
    text-decoration: underline;

    :hover {
        // text-decoration: underline;
        color: ${(props) => props.theme.colors.darker4};
    }
`;

export const Button = styled.button`
    border: none;
    background: ${(props) => props.theme.colors.primary3};
    padding: 10px 20px;
    font-size: 20px;
    cursor: pointer;

    :hover {
        background: ${(props) => props.theme.colors.darker3};
    }
`;

export const Row = styled.div`
    margin-top: 17px;
    margin-bottom: 17px;
`;

export const Cell = styled.div`
    display: inline-block;
    // width: 47.5%;
    width: 50%;

    > Input {
        width: 100%;
    }
`;

export const Page = styled.div`
    font-family: ${(props) => props.theme.font};
    background-color: ${(props) => props.theme.colors.lighter1};
    // padding: 5px 100px 100px 100px;
    font-size: 20px;
    height: 100vh;
    width: 100vw;
    display: table-cell;
    vertical-align: middle;
`;

export const Icon = styled.span`
    font-family: 'Material Icons', emoji;
    font-weight: normal;
    font-style: normal;
    font-size: 30px; /* Preferred icon size */
    display: inline-block;
    line-height: 1;
    text-transform: none;
    letter-spacing: normal;
    word-wrap: normal;
    white-space: nowrap;
    direction: ltr;
    vertical-align: middle;

    /* Support for all WebKit browsers. */
    -webkit-font-smoothing: antialiased;
    /* Support for Safari and Chrome. */
    text-rendering: optimizeLegibility;

    /* Support for Firefox. */
    -moz-osx-font-smoothing: grayscale;

    /* Support for IE. */
    font-feature-settings: 'liga';
`;

export function CustomIcon({ code, ...props }: { code: string }) {
    return (
        <Icon {...props}>
            <span className="material-icons-sharp">{code}</span>
        </Icon>
    );
}

export const DeleteIcon = (props: JSX.IntrinsicAttributes) => (
    <CustomIcon code="delete_outline" {...props} />
);
export const DownloadIcon = () => <CustomIcon code="file_download" />;
