<template>
  <div>
    <h4>Digite um comando:</h4>
    <InputText
      v-model="prompt"
      placeholder="Ex: crie um usuário chamado João com status ativo"

    />
    <Button
      label="Enviar"
      icon="pi pi-send"
      @click="enviarPrompt"
    />
    <p v-if="resposta">Resposta: {{ resposta }}</p>
  </div>
  <UserList />
</template>

<script setup lang="ts">
import { ref } from 'vue'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import UserList from './components/UsersList.vue'

const prompt = ref('')
const resposta = ref('')

const enviarPrompt = async () => {
  const res = await fetch('/api/prompt', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ prompt: prompt.value }),
  })
  resposta.value = await res.text()
}
</script>

<style scoped>

</style>
