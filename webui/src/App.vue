<template>
  <div class="container">
    <h2>Digite um comando:</h2>

    <InputText
      v-model="prompt"
      placeholder="Ex: crie um usuário chamado João com status ativo"
      class="p-inputtext-lg"
    />
    <Button
      label="Enviar"
      icon="pi pi-send"
      @click="enviarPrompt"
      class="p-button-lg p-button-outlined mt-3"
    />

    <p v-if="resposta" class="mt-3">Resposta: {{ resposta }}</p>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'

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
.container {
  max-width: 600px;
  margin: 40px auto;
  text-align: center;
  padding: 20px;
}
</style>
