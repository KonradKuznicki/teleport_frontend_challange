import { Page } from '../general/Elements';
import { PageHead } from '../Files/PageHead';
import { FilesList, FileStats } from '../Files/FileList';
import React, { useState } from 'react';

export function ListFilesPage({
    files,
    path,
}: {
    path: string[];
    files: FileStats[];
}) {
    const [search, onSearch] = useState('');

    console.log(search);
    return (
        <Page>
            <PageHead onSearch={onSearch} />
            <FilesList
                path={path}
                files={files.filter((i) => i.name.includes(search))}
            />
        </Page>
    );
}
