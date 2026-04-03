import React from 'react';
import {render} from '@testing-library/react';
import {Programmer} from './Programmer';

const mockProgrammer = {
    name: 'John Doe',
    title: 'Senior Developer',
    company: 'TechCorp',
    picture: 'https://example.com/avatar.jpg',
    skills: [
        {name: 'Go', importance: 3, icon: 'go'},
        {name: 'React', importance: 2, icon: 'react'}
    ]
};

test('renders programmer name', () => {
    const {getByText} = render(
        <Programmer programmer={mockProgrammer}
                    search="" updateSearch={() => {}}/>
    );
    expect(getByText('John Doe')).toBeInTheDocument();
});

test('renders title and company', () => {
    const {getByText} = render(
        <Programmer programmer={mockProgrammer}
                    search="" updateSearch={() => {}}/>
    );
    expect(getByText(/Senior Developer/)).toBeInTheDocument();
    expect(getByText(/TechCorp/)).toBeInTheDocument();
});

test('renders avatar image', () => {
    const {container} = render(
        <Programmer programmer={mockProgrammer}
                    search="" updateSearch={() => {}}/>
    );
    const img = container.querySelector('img');
    expect(img).toBeInTheDocument();
    expect(img.src).toBe('https://example.com/avatar.jpg');
});

test('renders skills list', () => {
    const {getByText} = render(
        <Programmer programmer={mockProgrammer}
                    search="" updateSearch={() => {}}/>
    );
    expect(getByText('Go')).toBeInTheDocument();
    expect(getByText('React')).toBeInTheDocument();
});

test('renders programmer without picture', () => {
    const noPicProgrammer = {...mockProgrammer, picture: null};
    const {getByText} = render(
        <Programmer programmer={noPicProgrammer}
                    search="" updateSearch={() => {}}/>
    );
    expect(getByText('John Doe')).toBeInTheDocument();
});
