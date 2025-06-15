<script setup>
import ToDoCard from "./ToDoCard.vue";
import todoData from "@/todos.json";
import { ref } from 'vue';

const todos = ref(todoData);

const removeTodo = (index) => {
    todos.value.splice(index, 1);
}

</script>

<template>
<section class="p-6">
    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4 lg:gap-6">
        <ToDoCard v-for="(todo, index) in todos" :key="todo.id" :todo="todo" :removeCallBack="(index) => removeTodo(index)"/>
    </div>
</section>

</template>