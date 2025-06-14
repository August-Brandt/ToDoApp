<script setup>
import Card from "@/components/Card.vue";
import { defineProps } from "vue";

defineProps({
    todo: Object,
});
</script>

<template>
    <Card>
        <h1 class="text-lg font-bold">{{ todo.title }}</h1>
        <p>
            {{ todo.description }}
        </p>
        <div class="border border-gray-300 mt-2"></div>
        <h3 class="text-sm text-gray-500">Finish by: {{ todo.doDate }}</h3>
    </Card>
</template>