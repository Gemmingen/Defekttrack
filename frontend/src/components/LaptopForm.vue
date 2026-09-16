<script setup lang="ts">
import { ref } from 'vue';
import { OnyxButton, OnyxInput, OnyxSelect, OnyxHeadline } from 'sit-onyx';

import { FehlerKategorie, type LaptopInput } from '../api/generated';

const emit = defineEmits<{ (e: 'submit', payload: LaptopInput): void }>();


const categories: { label: string; value: FehlerKategorie }[] = [
  { label: 'Display Schaden', value: FehlerKategorie.DISPLAY_SCHADEN },
  { label: 'Akku Defekt', value: FehlerKategorie.AKKU_DEFEKT },
  { label: 'Tastatur Defekt', value: FehlerKategorie.TASTATUR_DEFEKT },
  { label: 'Software / OS', value: FehlerKategorie.SOFTWARE_OS },
  { label: 'Kein Fehler', value: FehlerKategorie.KEIN_FEHLER }
];

const form = ref<LaptopInput>({
  marke: '',
  name: '',
  os: 'Windows 11 Enterprise',
  fehler: FehlerKategorie.DISPLAY_SCHADEN
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
      <OnyxInput label="Hersteller / Marke" v-model="form.marke" required placeholder="z.B. Lenovo" />
      <OnyxInput label="Gerätename / Modell" v-model="form.name" required placeholder="z.B. ThinkPad T14" />
      <OnyxInput label="Betriebssystem" v-model="form.os" required />
      
      <OnyxSelect 
        label="Fehlerkategorie" 
        listLabel="Fehlerkategorie auswählen"
        v-model="form.fehler" 
        :options="categories" 
      />

      <OnyxButton label="Gerät Registrieren" type="submit" color="primary" mode="default" />
    </form>
  </div>
</template>