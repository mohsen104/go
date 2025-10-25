const TodoList = [];

async function main() {
  const p1 = performance.now();
  for (let i = 1; i < 50; i++) {
    await GetTodo(i);
  }
  const p2 = performance.now();

  console.log(TodoList, (p2 - p1).toFixed(2) + "ms");
}

async function GetTodo(id) {
  await GetUrl(`https://jsonplaceholder.typicode.com/todos/${id}`);
}

async function GetUrl(url) {
  try {
    const response = await fetch(url);

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const resBody = await response.text();
    TodoList.push(resBody);
  } catch (error) {
    console.error("Error fetching data:", error);
    throw error; // Re-throw to maintain similar behavior to panic
  }
}

// Execute the main function
main().catch(console.error);
