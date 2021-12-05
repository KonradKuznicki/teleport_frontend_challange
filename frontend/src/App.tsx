import React from 'react';
import { ListFilesPage } from './Pages/ListFilesPage';
import { ThemeProvider } from 'styled-components';
import { theme } from './general/Elements';

function App() {
    return (
        <ThemeProvider theme={theme}>
            <ListFilesPage />
        </ThemeProvider>
    );
}

export default App;
