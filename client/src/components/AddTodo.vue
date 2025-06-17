<script setup>
</script>


<template>
    <section class="bg-green-300 flex justify-center border-t-2 border-green-100">
        <div class="w-sm md:w-lg p-3">
            <form class="grid grid-cols-1 md:grid-cols-3 gap-2">
                <div class="md:col-span-2 grid grid-rows-3 gap-2">
                    <input type="text" name="titel" class="bg-gray-100 rounded-lg px-2 h-6"/>
                    <textarea name="description" class="bg-gray-100 rounded-lg px-2 row-span-2 row-start-2 resize-none"></textarea>
                </div>
                <div class="flex align-center items-center justify-center">
                    <button type="submit" class="rounded-lg bg-gray-100 px-2 py-0.5 hover:bg-gray-300 hover:shadow">Add ToDo</button>
                </div>
            </form>
        </div>
    </section>
</template>