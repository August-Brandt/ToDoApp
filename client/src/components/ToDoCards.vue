<script setup>
import Card from "@/components/Card.vue"
import todoData from "@/todos.json"
import { ref } from 'vue'

const todos = ref(todoData);
console.log(todos.value);
</script>

<template>
<section class="p-6">
    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4 lg:gap-6">
        <Card v-for="todo in todos" :key="todo.id">
            <h1 class="text-lg font-bold">{{ todo.title }}</h1>
            <p>{{ todo.description }}</p>
        </Card>
    </div>
</section>

</template>