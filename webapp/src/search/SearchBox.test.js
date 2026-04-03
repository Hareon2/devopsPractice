import React from 'react';
import {render, fireEvent} from '@testing-library/react';
import {SearchBox} from './SearchBox';

test('renders input with placeholder', () => {
    const {getByPlaceholderText} = render(
        <SearchBox search="" updateSearch={() => {}}/>
    );
    const input = getByPlaceholderText(/type skill name to filter/i);
    expect(input).toBeInTheDocument();
});

test('displays current search value', () => {
    const {getByDisplayValue} = render(
        <SearchBox search="React" updateSearch={() => {}}/>
    );
    expect(getByDisplayValue('React')).toBeInTheDocument();
});

test('calls updateSearch on input change', () => {
    const mockUpdate = jest.fn();
    const {getByPlaceholderText} = render(
        <SearchBox search="" updateSearch={mockUpdate}/>
    );
    const input = getByPlaceholderText(/type skill name to filter/i);
    fireEvent.change(input, {target: {value: 'Go'}});
    expect(mockUpdate).toHaveBeenCalledWith('Go');
});

test('renders with empty search', () => {
    const {getByPlaceholderText} = render(
        <SearchBox search="" updateSearch={() => {}}/>
    );
    const input = getByPlaceholderText(/type skill name to filter/i);
    expect(input.value).toBe('');
});
