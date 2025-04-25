<template>
  <div>
    <h3>Usuários</h3>
    <DataTable :value="users" :paginator="false" :rows="5">
      <Column field="id" header="ID"></Column>
      <Column field="name" header="Nome"></Column>
      <Column field="status" header="Status"></Column>
      <Column field="created_at" header="Criado em"></Column>
    </DataTable>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'

interface User {
  id: number
  name: string
  status: string
  created_at: string
}

const users = ref<User[]>([])

onMounted(async () => {
  const response = await axios.get('/api/users')
  users.value = response.data
})
</script>
