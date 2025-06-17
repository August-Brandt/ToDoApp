<script setup>
import Card from "@/components/Card.vue";
import { defineProps, ref } from "vue";

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
        <div class="border border-gray-300 mt-2"></div>
        <h3 class="text-sm text-gray-500">Finish by: {{ todo.doDate }}</h3>
    </Card>
</template>