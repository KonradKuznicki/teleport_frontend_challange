import React from 'react';

function unit(magnitude: number) {
    switch (magnitude) {
        case 0:
            return '';
        case 1:
            return 'K';
        case 2:
            return 'M';
        case 3:
            return 'G';
        case 4:
            return 'T';
        case 5:
            return 'P';
        default:
            return '*10^' + magnitude; // (heh not needed, but let's be future proof ;) )
    }
}

export function format(size: number): string {
    const magnitude = Math.floor(Math.log(size) / Math.log(1024));
    if (magnitude) {
        return (
            Math.floor(size / Math.pow(1024, magnitude)) + unit(magnitude) + 'B'
        );
    }
    return size + 'B';
}

export function Size({ size, folder }: { size: number; folder?: boolean }) {
    if (folder) {
        return <span>--</span>;
    }
    return <span>{format(size)}</span>;
}
