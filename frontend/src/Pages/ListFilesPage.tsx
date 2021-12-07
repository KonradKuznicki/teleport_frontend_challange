import { Page, Title } from '../general/Elements';
import { PageHead } from '../Files/PageHead';
import { FilesList, FileStats } from '../Files/FileList';
import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import { useFetch } from '../API/useFetch';
import { Alert } from '../general/Alert';

export function ListFilesPage({
    files,
    path,
}: {
    path: string[];
    files: FileStats[];
}) {
    console.log(files, path);
    return (
        <Page>
            <PageHead />
            <FilesList path={path} files={files} />
        </Page>
    );
}

function NotFound() {
    return (
        <Page>
            <Alert type="danger">
                <Title>Page not found</Title>
                <Link to={'/files'}>Go back to known files</Link>
            </Alert>
        </Page>
    );
}
export function ListFilesRouteablePage() {
    const params = useLocation();
    const path = params.pathname;
    if (path.indexOf('/files') !== 0) {
        return <NotFound />;
    }
    return <ListFilesLoadablePage path={path.substring(7)} />;
}

export function ListFilesLoadablePage({ path }: { path: string }) {
    const { data, isLoading, isError } = useFetch<FileStats[]>(
        'https://localhost:3001/API/v1/files/' + path,
    );
    if (isError) {
        return (
            <Page>
                <Alert type="danger">{isError}</Alert>
            </Page>
        );
    }
    if (isLoading) {
        return (
            <Page>
                <Alert type="info">Please wait loading...</Alert>
            </Page>
        );
    }

    return (
        <ListFilesPage
            files={data as FileStats[]}
            path={path.split('/').filter((i) => i)}
        />
    );
}
