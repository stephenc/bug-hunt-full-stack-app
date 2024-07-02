import React, { useState, useEffect } from 'react';
import TodoList from './components/TodoList';
import './App.css';

function App() {
    const [todos, setTodos] = useState([]);
    const [text, setText] = useState('');

    useEffect(() => {
        fetch('/api/todos')
            .then(response => response.json())
            .then(data => setTodos(data));
    }, []);

    const addTodo = () => {
        fetch('/api/todos', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ text }),
        })
            .then(response => response.json())
            .then(newTodo => setTodos([...todos, newTodo]));
    };

    return (
        <div className="App">
            <h1>TODO App</h1>
            <input
                type="text"
                value={text}
                onChange={e => setText(e.target.value)}
            />
            <button onClick={addTodo}>Add TODO</button>
            <TodoList todos={todos} />
        </div>
    );
}

export default App;
