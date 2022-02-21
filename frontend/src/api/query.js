export const addTodoQuery = (title, description, priority) => {
    return {
        query: `
mutation addTodo($title: String!, $description: String!, $priority: Int!) {
  addTodo(input: 
  {
      title: $title, 
      description: $description,
      priority: $priority
   }) {
    id
    title
    description
    priority
    createdAt
  }
}
        `,
        variables: {
            "title": title,
            "description": description,
            "priority": priority,
        }

    }
}
export const updateTodoQuery = (id, title, description, priority) => {
    return {
        query: `
mutation updateTodo($id: Int!,$title: String!, $description: String!, $priority: Int!) {
  updateTodo(input: {id: $id,title: $title, description: $description, priority: $priority}) {
    id
    title
    description
    priority
    createdAt
  }
}
        `,
        variables: {
            id: id,
            title: title,
            description: description,
            priority: priority,
        }
        ,
    }
}


export const getTodoQuery = (sortProperty, orderOption) => {
    return {
        query: `
query todos($sortProperty: SortOptions,$orderOption: OrderOptions) {
  todos(sortProperty:$sortProperty,orderOption:$orderOption){
    id
    title
    description
    priority
    createdAt
  }
}
        `,
        variables:
            {
                sortProperty: sortProperty,
                orderOption: orderOption,
            }
        ,
    }
}

export const removeTodoQuery = (id) => {
    return {
        query: `
mutation removeTodo($id: Int!) {
  removeTodo(input: {id: $id}) {
    id
    title
    description
    priority
    createdAt
  }
}
        `,
        variables: {
            id: id
        }
        ,
    }
}
