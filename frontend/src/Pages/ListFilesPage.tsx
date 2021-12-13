import { Page } from '../general/Elements';
import { PageHead } from '../Files/PageHead';
import { FilesList, FileStats } from '../Files/FileList';
import React, { useState } from 'react';
import { format } from '../Files/Size';

function compare(a: any, b: any) {
    if (typeof a === 'string') {
        return a.localeCompare(b);
    }
    return a - b;
}

export function ListFilesPage({
    files,
    path,
}: {
    path: string[];
    files: FileStats[];
}) {
    const [search, onSearch] = useState('');
    const [sortBy, setSortBy] = useState('');

    const ls = search.toLowerCase();
    const filteredFiles = files.filter(
        (i) =>
            i.name.toLowerCase().includes(ls) ||
            i.type.toLowerCase().includes(ls) ||
            (i.type === 'folder' ? '--' : format(i.size))
                .toLowerCase()
                .includes(ls),
    );
    if (sortBy) {
        filteredFiles.sort((a: any, b: any) => compare(a[sortBy], b[sortBy]));
    }

    return (
        <Page>
            <PageHead onSearch={onSearch} />
            <FilesList
                path={path}
                files={filteredFiles}
                sortBy={sortBy}
                onSort={setSortBy}
            />
        </Page>
    );
}
