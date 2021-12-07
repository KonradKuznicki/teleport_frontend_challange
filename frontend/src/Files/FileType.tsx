import { CustomIcon } from '../general/Elements';
import React from 'react';

function type2icon(type: string) {
    switch (type.toLowerCase()) {
        case 'png':
        case 'gif':
        case 'jpg':
            return 'image';
        case 'folder':
            return 'folder_open';
        case 'txt':
        case 'pdf':
            return 'description';
        default:
            return '?';
    }
}

export function FileType({ type }: { type: string }) {
    return <CustomIcon code={type2icon(type)} />;
}