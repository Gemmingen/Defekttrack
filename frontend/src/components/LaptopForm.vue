<script setup lang="ts">
import { ref } from 'vue';
import { OnyxButton, OnyxInput, OnyxSelect, OnyxHeadline } from 'sit-onyx';
import type { LaptopInput, FehlerKategorie } from '../types/laptop';

const emit = defineEmits<{ (e: 'submit', payload: LaptopInput): void }>();

const categories: FehlerKategorie[] = ['Hardware', 'Display', 'Akku', 'Software', 'Netzwerk', 'Sonstiges'];

const form = ref<LaptopInput>({
  marke: '',
  name: '',
  os: 'Windows 11 Enterprise',
  fehler: 'Hardware'
});

const handleSubmit = () => {
  if (!form.value.marke || !form.value.name) return;
  emit('submit', { ...form.value });
  form.value.marke = '';
  form.value.name = '';
};
</script>

<template>
  <div class="sd-card">
    <div class="card-header-badge">SCHWARZ DIGITS SUPPORT</div>
    <OnyxHeadline is="h3">Hardware-Defekt melden</OnyxHeadline>
    
    <form @submit.prevent="handleSubmit" class="form-grid">
      <OnyxInput label="Hersteller / Marke" v-model="form.marke" required placeholder="z.B. Lenovo" message= "Erlaubte Marken: HP, Lenovo, Dell, Apple, Asus, Acer"/>
      <OnyxInput label="Gerätename / Modell" v-model="form.name" required placeholder="z.B. ThinkPad T14" />
      <OnyxInput label="Betriebssystem" v-model="form.os" required />
      
      <OnyxSelect 
        label="Fehlerkategorie" 
        listLabel="Fehlerkategorie auswählen"
        v-model="form.fehler" 
        :options="categories.map(c => ({ label: c, value: c }))" 
      />

      <OnyxButton label="Gerät Registrieren" type="submit" color="primary" mode="default" />
    </form>
  </div>
</template>

<style scoped>
.sd-card {
  background: var(--onyx-color-base-background-blank);
  border: 1px solid var(--onyx-color-base-neutral-200);
  border-radius: var(--onyx-radius-md);
  padding: var(--onyx-spacing-xl);
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.04);
}

.card-header-badge {
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 1px;
  color: #00a3e0;
  margin-bottom: var(--onyx-spacing-2xs);
}

.form-grid {
  display: flex;
  flex-direction: column;
  gap: var(--onyx-spacing-md);
  margin-top: var(--onyx-spacing-md);
}
</style>