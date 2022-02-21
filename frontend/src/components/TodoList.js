import React, {useState, useEffect} from "react";
import TodoForm from "./TodoForm";
import Todo from "./Todo";
import {post} from "../api/api";
import {addTodoQuery, getTodoQuery, removeTodoQuery, updateTodoQuery} from "../api/query";

function TodoList() {
    const [todos, setTodos] = useState([]);
    useEffect(() => {
        LoadTodos()
    },[])
    const addTodo = async function (todo) {
        if (!todo.text || /^\s*$/.test(todo.text)) {
            return;
        }
        console.log("todo add",todo)
        const response = await post(null, addTodoQuery(todo.text, todo.text, 8))
        console.log(response)
        const newTodos = [response.data.data.addTodo, ...todos];

        setTodos(newTodos);
        console.log(...todos);
    }

    const updateTodo = async (id, newValue) => {
        if (!newValue.text || /^\s*$/.test(newValue.text)) {
            return;
        }
console.log("newval",newValue )
        const response = await post(null, updateTodoQuery(id,newValue.text, newValue.text, 8))
        console.log(response)
        await LoadTodos()
        // setTodos((prev) =>
        //     prev.map((item) => (item.id === todoId ? newValue : item))
        // );
    };

    const removeTodo = async (id) => {
        const response = await post(null, removeTodoQuery(id))
        console.log(response)
        await LoadTodos()
        // setTodos(removedArr);
    };

    const completeTodo = function* (id) {
        const response = yield post(null, getTodoQuery("createdAt", "asc"))
        console.log("---", response)
        let updatedTodos = todos.map((todo) => {
            if (todo.id === id) {
                todo.isComplete = !todo.isComplete;
            }
            return todo;
        });
        setTodos(updatedTodos);
    };
    async function LoadTodos() {
        const response = await post(null, getTodoQuery("createdAt", "asc"))
        setTodos(response.data.data.todos)
    }
    return (
        <>
            <h1>TODO LIST</h1>
            <TodoForm onSubmit={addTodo}/>
            <Todo
                todos={todos}
                completeTodo={completeTodo}
                removeTodo={removeTodo}
                updateTodo={updateTodo}
            />
        </>
    );
}

export default TodoList;
