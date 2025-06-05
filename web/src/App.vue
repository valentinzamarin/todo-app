<script setup>
import { ref } from 'vue'

const title = ref('')

const addNewTask = async (e) => {
  if (title.value.trim() === "") {
    alert('Empty field')
    return
  }


  const response = await fetch('/api/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      title: title.value
    })
  })

  if (response.ok) {
    const newTask = await response.json()
    console.log('Задача создана:', newTask)
    title.value = ''
  } else {
    const errorText = await response.text()
    console.error('Server error:', errorText)
    alert('Err')
  }

}
</script>

<template>
  <div class="h-screen w-full flex items-center justify-center bg-[#e5e7eb]">
    <div class="bg-white rounded shadow p-6 m-4 w-full lg:w-3/4 lg:max-w-lg">
      <form @submit.prevent="addNewTask">
        <input v-model="title"
          class="w-full bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded-md px-3 py-2 transition duration-300 ease focus:outline-none focus:border-slate-400 hover:border-slate-300 shadow-sm focus:shadow"
          type="text" placeholder="Create task.." />
      </form>
    </div>
  </div>
</template>