import React from 'react';
import { ListFilesRouteablePage } from './Pages/ListFilesPage';
import { ThemeProvider } from 'styled-components';
import { theme, Title } from './general/Elements';
import { BrowserRouter, Routes, Route, Link } from 'react-router-dom';

function NotFound() {
    return (
        <div>
            <Title>Page not found</Title>
            <Link to={'/files'}>Go back to known files</Link>
        </div>
    );
}

function App() {
    return (
        <ThemeProvider theme={theme}>
            <BrowserRouter>
                <Routes>
                    <Route
                        path="/files/:path"
                        element={<ListFilesRouteablePage />}
                    />
                    <Route path="/files" element={<ListFilesRouteablePage />} />
                    <Route path="*" element={<NotFound />} />
                </Routes>
            </BrowserRouter>
        </ThemeProvider>
    );
}

export default App;
