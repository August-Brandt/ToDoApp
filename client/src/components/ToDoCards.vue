<script setup>
import Card from "@/components/Card.vue"

</script>

<template>
<section class="p-6">
    <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4 lg:gap-6">
        <Card>
            <h2 class="text-lg font-bold">Todo Title 1</h2>
            <p>
                Here is a discription of what the todo is about.
            </p>
        </Card>
        <Card>
            <h2 class="text-lg font-bold">Todo Title 2</h2>
            <p>
                Here is a discription of what the todo is about.
            </p>
        </Card>
        <Card>
            <h2 class="text-lg font-bold">Todo Title 3</h2>
            <p>
                Here is a discription of what the todo is about.
            </p>
        </Card>
        <Card>
            <h2 class="text-lg font-bold">Todo Title 4</h2>
            <p>
                Here is a discription of what the todo is about.
            </p>
        </Card>
    </div>
</section>

</template>