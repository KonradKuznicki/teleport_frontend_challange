import { Page } from '../general/Elements';
import { PageHead } from '../Files/PageHead';
import { FilesList, FileStats } from '../Files/FileList';
import React from 'react';
import { useParams } from 'react-router-dom';
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

export function ListFilesRouteablePage() {
    const params = useParams();
    const path = params.path || '';
    const { data, isLoading, isError } = useFetch<FileStats[]>(
        'https://localhost:3001/API/v1/files' + path,
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
