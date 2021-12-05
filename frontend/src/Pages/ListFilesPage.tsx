import { Page } from '../general/Elements';
import { PageHead } from '../Files/PageHead';
import { FilesList } from '../Files/FileList';
import React from 'react';

export function ListFilesPage() {
    return (
        <Page>
            <PageHead />
            <FilesList />
        </Page>
    );
}
