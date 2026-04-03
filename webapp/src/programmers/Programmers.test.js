import React from 'react';
import {render} from '@testing-library/react';
import {Programmers} from './Programmers';

const mockProgrammers = [
    {
        name: 'Alice',
        title: 'Backend Dev',
        company: 'AlphaCo',
        picture: 'https://example.com/alice.jpg',
        skills: [{name: 'Go', importance: 3, icon: 'go'}]
    },
    {
        name: 'Bob',
        title: 'Frontend Dev',
        company: 'BetaInc',
        picture: 'https://example.com/bob.jpg',
        skills: [{name: 'React', importance: 2, icon: 'react'}]
    }
];

test('renders list of programmers', () => {
    const {getByText} = render(
        <Programmers programmers={mockProgrammers}
                     search="" updateSearch={() => {}}/>
    );
    expect(getByText('Alice')).toBeInTheDocument();
    expect(getByText('Bob')).toBeInTheDocument();
});

test('renders empty list without errors', () => {
    const {container} = render(
        <Programmers programmers={[]}
                     search="" updateSearch={() => {}}/>
    );
    // Should render a div with no children
    expect(container.firstChild).toBeInTheDocument();
    expect(container.firstChild.children.length).toBe(0);
});

test('renders single programmer', () => {
    const {getByText, queryByText} = render(
        <Programmers programmers={[mockProgrammers[0]]}
                     search="" updateSearch={() => {}}/>
    );
    expect(getByText('Alice')).toBeInTheDocument();
    expect(queryByText('Bob')).toBeNull();
});
