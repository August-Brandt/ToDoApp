<script setup>
import { ref } from 'vue';
import Navbar from '@/components/Navbar.vue'

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
  <h1>Vue ToDo</h1>
  <h3>Hello {{ name }}</h3>
  <button @click="clicking">Clicked {{ clicks }} times</button>
  <form @submit.prevent="addTodo">
    <label for="newTodo">Add ToDo</label>
    <input type="text" id="newTodo" name="newTodo" v-model="newTodo">
    <button type="submit">Add Todo</button>
  </form>
  <h3 v-if="tasks.length > 0" class="ml-3 underline text-xl">ToDos:</h3>
  <ul class="ml-5">
    <li v-for="(task, index) in tasks" :key="task"> 
      <button @click="removeTodo(index)" class="mr-2 bg-gray-300 center text-red-600 border-2 rounded-sm text-center w-5 h-5 font-bold flex justify-center">
        <i class="pi pi-times align-middle"></i>
      </button>
      <span class="font-bold underline">
        {{ task }}
      </span>
    </li>
  </ul>
</template>