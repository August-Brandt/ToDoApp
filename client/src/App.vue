<script setup>
import { ref, onMounted } from 'vue';
import Navbar from '@/components/Navbar.vue'
import Hero from '@/components/Hero.vue'
import ToDoCards from './components/ToDoCards.vue'
import AddTodo from './components/AddTodo.vue';
import axios from 'axios';


const todos = ref([]);

const addTodo = async (newTodo) => {
    try {
        const response = await axios.post("/api/addtodo", 
            JSON.stringify({
                id: "",
                title: newTodo.title,
                description: newTodo.description,
                doDate: newTodo.doDate,
                finished: false,
            }));
        todos.value.push(response.data);
    } catch (error) {
        console.error("Error creating new todo", error);
    }
};

const removeTodo = async (id) => {
    try {
        const response = await axios.post("/api/removetodo/" + id);
        todos.value = todos.value.filter((todo) => {
            return todo.id != id;
        });
    } catch (error) {
        console.error("Error creating new todo", error);
    }
}

const toggleFinished = async (id) => {
    const todo = todos.value.find(t => t.id === id);
    if (todo) {
        if (!todo.finished) {
            try {
                const response = await axios.post("/api/finishtodo/" + id);
                todo.finished = true;
            } catch (error) {
                console.error("Error finishing todo", error);
            }
        } else {
            try {
                const response = await axios.post("/api/unfinishtodo/" + id);
                todo.finished = false;
            } catch (error) {
                console.error("Error unfinishing todo", error);
            }
        }
    } else {
        console.Error("Unable to find todo with that id", todo);
    }
};

onMounted(async () => {
    try {
        const response = await axios.get("1/api/todos");
        todos.value = response.data;
    } catch (error) {
        console.error("Error fetching todos", error);
    }
})

</script>

<template>
    <Navbar/>
    <Hero/>
    <AddTodo @add="addTodo"/>
    <ToDoCards :todos="todos" @remove="removeTodo" @toggle="toggleFinished"/>
</template>