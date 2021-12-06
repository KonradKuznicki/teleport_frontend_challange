import React from 'react';
import { render, screen } from '@testing-library/react';
import { format, Size } from './Size';

test('renders 4k', () => {
    render(<Size size={4555} />);
    const size = screen.getByText(/4KB/i);
    expect(size).toBeInTheDocument();
});

test('renders 4 items', () => {
    render(<Size size={4} folder />);
    const size = screen.getByText(/4 items/i);
    expect(size).toBeInTheDocument();
});

test('formatter should deal with bytes', () => {
    expect(format(12)).toBe('12B');
});

test('formatter should deal with K,M,G,T,P bytes', () => {
    expect(format(1200)).toBe('1KB');
    expect(format(1200000)).toBe('1MB');
    expect(format(1200000000)).toBe('1GB');
    expect(format(1200000000000)).toBe('1TB');
    expect(format(1200000000000000)).toBe('1PB');
});

test('formatter should deal with extra large', () => {
    expect(format(Math.pow(10, 100))).toBe('4*10^33B');
});
