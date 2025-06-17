<script setup>
import { ref } from 'vue';
import Navbar from '@/components/Navbar.vue'
import Hero from '@/components/Hero.vue'
import ToDoCards from './components/ToDoCards.vue'
import AddTodo from './components/AddTodo.vue';
import todoData from "@/todos.json";
import { v4 } from 'uuid';


const todos = ref(todoData);

const addTodo = (newTodo) => {
    todos.value.push({
        id: v4(),
        finished: false,
        ...newTodo,
    });
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

</script>

<template>
    <Navbar/>
    <Hero/>
    <AddTodo @add="addTodo"/>
    <ToDoCards :todos="todos" @remove="removeTodo" @toggle="toggleFinished"/>
</template>