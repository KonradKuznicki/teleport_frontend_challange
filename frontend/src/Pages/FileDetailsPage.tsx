import { Page } from '../general/Elements';
import { FileDetailsPageHead } from '../Files/FileDetailsPageHead';
import { FileDetails } from '../Files/FileDetails';
import React from 'react';
import { FileStats } from '../Files/FileList';

export function FileDetailsPage({
    pathParts,
    details,
    ...props
}: {
    pathParts: string[];
    details: FileStats;
}) {
    return (
        <Page style={{ display: 'table-cell' }} {...props}>
            <FileDetailsPageHead />
            <FileDetails pathParts={pathParts} details={details} />
        </Page>
    );
}
