<script setup>
import ToDoCard from "./ToDoCard.vue";
import todoData from "@/todos.json";
import { ref, computed } from 'vue';

const todos = ref(todoData);

const removeTodo = (id) => {
    todos.value = todos.value.filter((todo) => {
        return todo.id != id;
    });
}
 
const sortedTodos = computed(() => {
    todos.value.sort(( a, b ) => {
        if (a.finished && b.finished) {
            return 0;
        }
        if (a.finished && !b.finished) {
            return 1;
        }
        return -1;
    });
    console.log(todos);
    return todos.value;
});

</script>

<template>
<section class="p-6">
    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4 lg:gap-6">
        <ToDoCard v-for="todo in sortedTodos" :key="todo.id" :todo="todo" :removeCallBack="() => removeTodo(todo.id)"/>
    </div>
</section>

</template>