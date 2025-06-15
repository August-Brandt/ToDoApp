<script setup>
import Card from "@/components/Card.vue";
import { defineProps } from "vue";

defineProps({
    todo: Object,
    removeCallBack: Function,
});
</script>

<template>
    <Card>
        <div class="grid grid-cols-5">
            <h1 class="text-lg font-bold col-span-4">{{ todo.title }}</h1>
            <button class="bg-gray-300 col-start-5 w-7 rounded-lg" @click="removeCallBack">
                <img src="@/assets/img/TrashCan.png" alt="">
            </button>
        </div>
        <p>
            {{ todo.description }}
        </p>
        <div class="border border-gray-300 mt-2"></div>
        <h3 class="text-sm text-gray-500">Finish by: {{ todo.doDate }}</h3>
    </Card>
</template>