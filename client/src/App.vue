<script setup>
import { ref } from 'vue';
import Navbar from '@/components/Navbar.vue'
import Hero from '@/components/Hero.vue'

const name = "August";
const clicks = ref(0);
const tasks = ref(["Make dinner", "Play CS with Theis", "Go to sleep"]);
const newTodo = ref('');

const clicking = () => { 
  console.log("Button clicked");
  clicks.value++;
};

const addTodo = () => {
  if (newTodo.value.trim() !== '') {
    tasks.value.push(newTodo.value.trim());
  }
  newTodo.value = '';
};

const removeTodo = (index) => {
  tasks.value.splice(index, 1);
}
</script>

<template>
  <Navbar/>
  <Hero/>
  
</template>