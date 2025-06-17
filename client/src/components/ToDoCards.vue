<script setup>
import ToDoCard from "./ToDoCard.vue";

import { defineProps, defineEmits, computed } from 'vue';

const props = defineProps({
    todos: Array,
});

const emit = defineEmits(['remove', 'toggle']);
 
const sortedTodos = computed(() => {
    props.todos.sort(( a, b ) => {
        if (a.finished && b.finished) {
            return 0;
        }
        if (a.finished && !b.finished) {
            return 1;
        }
        return -1;
    });
    return props.todos;
});

</script>

<template>
<section class="p-6">
    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4 lg:gap-6">
        <ToDoCard v-for="todo in sortedTodos" :key="todo.id" :todo="todo" :removeCallBack="() => emit('remove', todo.id)" @click="emit('toggle', todo.id)"/>
    </div>
</section>

</template>