<script setup>
import Card from "@/components/Card.vue";
import { defineProps, ref } from "vue";

const isToday = (dateString) => {
    const date = new Date(Date.parse(dateString));
    const dateNow = new Date(Date.now());
    if (date.getFullYear() == dateNow.getFullYear() && 
        date.getMonth() == dateNow.getMonth() && 
        date.getDate() == dateNow.getDate()) 
    {
        return true;    
    } else {
        return false;
    }
}

defineProps({
    todo: Object,
    removeCallBack: Function,
});
</script>

<template>
    <Card>
        <div class="grid grid-cols-5">
            <h1 v-if="todo.finished" class="text-lg font-bold col-span-4 line-through text-gray-400">
                {{ todo.title }}
            </h1>
            <h1 v-else class="text-lg font-bold col-span-4">
                {{ todo.title }}
            </h1>
            <div class=" col-start-5 text-right">
                <button class="bg-gray-300 w-7 rounded-lg hover:bg-gray-400 hover:shadow" @click="removeCallBack">
                    <img src="@/assets/img/TrashCan.png" alt="Remove button">
                </button>
            </div>
        </div>
        <p v-if="todo.finished" class="line-through text-gray-500">
            {{ todo.description }}
        </p>
        <p v-else>
            {{ todo.description }}
        </p>
        <div v-if="todo.doDate">
            <div class="border border-gray-300 mt-2"></div>
            <div v-if="todo.finished">
                <h3 class="line-through text-sm text-gray-400">Finish by: {{ todo.doDate }}</h3>
            </div>
            <div v-else>
                <h3 v-if="isToday(todo.doDate)" class="text-sm text-gray-900">Finish by: {{ todo.doDate }}</h3>
                <h3 v-else-if="Date.parse(todo.doDate) < Date.now()" class="text-sm font-bold text-red-400">Finish by: {{ todo.doDate }}</h3>
                <h3 v-else class="text-sm text-gray-400">Finish by: {{ todo.doDate }}</h3>
            </div>
        </div>
    </Card>
</template>