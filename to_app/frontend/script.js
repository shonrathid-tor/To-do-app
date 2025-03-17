const API_URL = "http://localhost:9999/todos";

async function fetchTasks() {
    const response = await fetch(API_URL);
    const tasks = await response.json();
    const todoList = document.getElementById("todoList");
    todoList.innerHTML = "";

    tasks.forEach(task => {
        const li = document.createElement("li");
        li.textContent = task.task;

        const deleteButton = document.createElement("button");
        deleteButton.textContent = "X";

        deleteButton.onclick = () => deleteTask(task.id);

        li.appendChild(deleteButton);
        
        todoList.appendChild(li);
    });
}

async function addTask() {
    const taskInput = document.getElementById("taskInput");
    const task = taskInput.value.trim();
    if (task === "") return;

    await fetch(API_URL, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ task, done: false }),
    });

    taskInput.value = "";
    fetchTasks();
}

async function deleteTask(id) {
    await fetch(`${API_URL}/${id}`, { method: "DELETE" });
    fetchTasks();
}

// Load tasks on page load
document.addEventListener("DOMContentLoaded", fetchTasks);
