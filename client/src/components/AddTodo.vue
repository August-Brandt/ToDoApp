<script setup>
import { ref, defineEmits } from 'vue';

const emit = defineEmits(['add']);

const title = ref('');
const description = ref('');

const addTodo = () => {
    emit('add', {
        title: title.value, 
        description: description.value,
        doDate: "2025-6-20",
    });
    title.value = "";
    description.value = "";
}

</script>


<template>
    <section class="bg-green-300 flex justify-center border-t-2 border-green-100">
        <div class="w-sm md:w-lg p-3">
            <form @submit.prevent="addTodo" class="grid grid-cols-1 md:grid-cols-3 gap-2">
                <div class="md:col-span-2 grid grid-rows-3 gap-2">
                    <input type="text" name="titel" v-model="title" class="bg-gray-100 rounded-lg px-2 h-6"/>
                    <textarea name="description" v-model="description" class="bg-gray-100 rounded-lg px-2 row-span-2 row-start-2 resize-none"></textarea>
                </div>
                <div class="flex align-center items-center justify-center">
                    <button type="submit" class="rounded-lg bg-gray-100 px-2 py-0.5 hover:bg-gray-300 hover:shadow">Add ToDo</button>
                </div>
            </form>
        </div>
    </section>
</template>