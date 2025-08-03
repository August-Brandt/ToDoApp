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
        const response = await axios.post("http://localhost:8080/api/newtodo", 
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

const removeTodo = (id) => {
    todos.value = todos.value.filter((todo) => {
        return todo.id != id;
    });
}

const toggleFinished = (id) => {
    const todo = todos.value.find(t => t.id === id);
    if (todo) {
        todo.finished = !todo.finished;
    }
};

onMounted(async () => {
    try {
        const response = await axios.get("http://localhost:8080/api/todos");
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