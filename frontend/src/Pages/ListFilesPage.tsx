import { Page, Title } from '../general/Elements';
import { PageHead } from '../Files/PageHead';
import { FilesList, FileStats } from '../Files/FileList';
import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import { useFetch } from '../API/useFetch';
import { Alert } from '../general/Alert';
import { FileDetailsPage } from './FileDetailsPage';

export function ListFilesPage({
    files,
    path,
}: {
    path: string[];
    files: FileStats[];
}) {
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

function maybeFileDetails(
    data: FileStats[] | undefined,
    path: string,
): FileStats | false {
    if (
        data &&
        data.length == 1 &&
        decodeURI(path).lastIndexOf(data[0].name) ==
            decodeURI(path).length - data[0].name.length &&
        data[0].type !== 'folder'
    ) {
        return data[0];
    }
    return false;
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

    const details = maybeFileDetails(data, path);
    if (details) {
        return (
            <FileDetailsPage details={details} pathParts={path.split('/')} />
        );
    }

    return (
        <ListFilesPage
            files={data as FileStats[]}
            path={path.split('/').filter((i) => i)}
        />
    );
}
