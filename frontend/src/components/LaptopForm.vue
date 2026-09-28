<script setup lang="ts">
import { ref } from 'vue';
import { OnyxButton, OnyxInput, OnyxSelect, OnyxHeadline } from 'sit-onyx';
import { FehlerKategorie, type LaptopInput } from '../api/generated';
import { CATEGORY_OPTIONS } from '../constants/categories';

const emit = defineEmits<{ (e: 'submit', payload: LaptopInput): void }>();

const form = ref<LaptopInput>({
  marke: '',
  name: '',
  os: 'Windows 11 Enterprise',
  fehler: FehlerKategorie.DISPLAY_SCHADEN,
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
      <OnyxInput
        label="Hersteller / Marke"
        v-model="form.marke"
        required
        placeholder="z.B. Lenovo"
        message="Erlaubte Marken: HP, Lenovo, Dell, Apple, Asus, Acer"
      />
      <OnyxInput
        label="Gerätename / Modell"
        v-model="form.name"
        required
        placeholder="z.B. ThinkPad T14"
        message="Exaktes Modell angeben"
      />
      <OnyxInput label="Betriebssystem" v-model="form.os" required />

      <OnyxSelect
        label="Fehlerkategorie"
        listLabel="Fehlerkategorie auswählen"
        v-model="form.fehler"
        :options="CATEGORY_OPTIONS"
      />
      <OnyxButton label="Gerät Registrieren" type="submit" color="primary" mode="default" />
    </form>
  </div>
</template>

<style scoped>
.sd-card {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: var(--onyx-radius-md);
  padding: var(--onyx-spacing-lg);
}
.card-header-badge {
  font-size: 10px;
  font-weight: 800;
  letter-spacing: 1px;
  color: #00a3e0;
  margin-bottom: var(--onyx-spacing-2xs);
}
.form-grid {
  display: flex;
  flex-direction: column;
  gap: var(--onyx-spacing-sm);
  margin-top: var(--onyx-spacing-md);
}
</style>