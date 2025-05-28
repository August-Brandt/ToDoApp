<script setup>
import logo from '@/assets/img/logo.png'

</script>

<template>
  <nav class="bg-green-400 border-b-2 border-green-100">
    <div class="mx-auto max-w-7xl px-2 sm:px-6 lg:px-8">
      <div class="flex h-15 items-center justify-between">
        <div class="flex flex-1 items-center justify-center md:justify-start">
          <!-- Logo -->
          <a href="/" class="flex flex-shrink-0 items-center mr-4">
            <img :src="logo" alt="logo" class="h-10 w-auto bg-white rounded-2xl">
            <span class="hidden md:block text-white text-xl font-extrabold ml-2">
              Vue ToDo
            </span>
          </a>
          
          <div class="">
            <div class="flex space-x-2">
              <a href="/" class="text-white text-l font-bold bg-green-600 py-1 px-2 rounded-lg hover:bg-green-800">
                Home
              </a>
  
              <a href="/todos" class="text-white text-l font-bold py-1 px-2 rounded-lg hover:bg-green-800">
                ToDo List
              </a>
  
              <a href="/about" class="text-white text-l font-bold  py-1 px-2 rounded-lg hover:bg-green-800">
                About
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </nav> 
</template>