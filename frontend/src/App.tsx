import React from 'react';
import { ThemeProvider } from 'styled-components';
import { theme } from './general/Elements';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { RouteablePage } from './Pages/RouteablePage';

function App() {
    return (
        <ThemeProvider theme={theme}>
            <BrowserRouter>
                <Routes>
                    <Route path="*" element={<RouteablePage />} />
                </Routes>
            </BrowserRouter>
        </ThemeProvider>
    );
}

export default App;
