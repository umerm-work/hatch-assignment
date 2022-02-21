import React, { useState, useEffect, useRef } from 'react';

function TodoForm(props) {
  const [inputTitle, setInputTitleTitle] = useState(props.edit ? props.edit.value : '');
  const [inputDescription, setInputTitleDescription] = useState(props.edit ? props.edit.value : '');

  const inputRefTitle = useRef(null);
  const inputRefDescription = useRef(null);

  useEffect(() => {
    // inputRefTitle.current.focus();
    // inputRefDescription.current.focus();
  });

  const handleTitleChange = e => {
    setInputTitleTitle(e.target.value);
  };

  const handleDescriptionChange = e => {
    setInputTitleDescription(e.target.value);
  };

  const handleSubmit = e => {
    e.preventDefault();

    console.log("handleSubmit",e)
    props.onSubmit({
      id: Math.floor(Math.random() * 10000),
      text: inputTitle
    });
    setInputTitleTitle('');
    setInputTitleDescription('');
  };

  return (
    <form onSubmit={handleSubmit} className='todo-form'>
      {props.edit ? (
        <>
          <input
            placeholder='Update your title'
            value={inputTitle}
            onChange={handleTitleChange}
            name='title'
            ref={inputRefTitle}
            className='todo-input edit'
          />
          <input
            placeholder='Update your description'
            value={inputDescription}
            onChange={handleDescriptionChange}
            name='description'
            ref={inputRefDescription}
            className='todo-input edit'
          />
          <button onClick={handleSubmit} className='todo-button edit'>
            Update
          </button>
        </>
      ) : (
        <>
          <input
            placeholder='Add a Title'
            value={inputTitle}
            onChange={handleTitleChange}
            name='title'
            className='todo-input'
            ref={inputRefTitle}
          />
          <input
            placeholder='Add a description'
            value={inputDescription}
            onChange={handleDescriptionChange}
            name='description'
            className='todo-input'
            ref={inputRefDescription}
          />
          <button onClick={handleSubmit} className='todo-button'>
            Add
          </button>
        </>
      )}
    </form>
  );
}

export default TodoForm;