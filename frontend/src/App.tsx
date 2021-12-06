import React from 'react';
import { ListFilesRouteablePage } from './Pages/ListFilesPage';
import { ThemeProvider } from 'styled-components';
import { theme } from './general/Elements';
import { BrowserRouter, Routes, Route } from 'react-router-dom';

function App() {
    return (
        <ThemeProvider theme={theme}>
            <BrowserRouter>
                <Routes>
                    <Route path="*" element={<ListFilesRouteablePage />} />
                </Routes>
            </BrowserRouter>
        </ThemeProvider>
    );
}

export default App;
