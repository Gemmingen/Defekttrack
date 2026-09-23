<script setup lang="ts">
import { ref } from 'vue';
import { OnyxHeadline, OnyxButton, OnyxInput } from 'sit-onyx';
import { iconTrash, iconPlus } from '@sit-onyx/icons';
import { CATEGORY_LABELS } from '../constants/categories';
import type { Laptop } from '../api/generated';

const props = defineProps<{
  selectedLaptop: Laptop | null;
}>();

const emit = defineEmits<{
  (e: 'removeLaptop', id: number): void;
  (e: 'addLog', payload: { bearbeiter: string; notiz: string }): void;
  (e: 'removeLog', logId: number): void;
}>();

const bearbeiter = ref('');
const notiz = ref('');

const handleCreateLog = () => {
  if (!props.selectedLaptop || !bearbeiter.value || !notiz.value) return;
  emit('addLog', { bearbeiter: bearbeiter.value, notiz: notiz.value });
  notiz.value = '';
};
</script>

<template>
  <section class="detail-panel">
    <div v-if="selectedLaptop" class="sd-card detail-card">
      <div class="detail-header">
        <div>
          <span class="device-id">ID #{{ selectedLaptop.id }}</span>
          <OnyxHeadline is="h2">{{ selectedLaptop.marke }} {{ selectedLaptop.name }}</OnyxHeadline>
          <div class="device-specs">
            <span><strong>Kategorie:</strong> {{ CATEGORY_LABELS[selectedLaptop.fehler] || selectedLaptop.fehler }}</span>
            <span><strong>OS:</strong> {{ selectedLaptop.os }}</span>
          </div>
        </div>
        <OnyxButton
          label="Löschen"
          color="danger"
          mode="outline"
          :icon="iconTrash"
          @click="emit('removeLaptop', selectedLaptop.id)"
        />
      </div>
      <div class="sd-divider"></div>
      <!-- Support Logs -->
      <OnyxHeadline is="h3">IT-Support Logs</OnyxHeadline>

      <div class="logs-wrapper">
        <div v-for="log in selectedLaptop.it_logs" :key="log.id" class="log-card">
          <div class="log-header">
            <span class="log-author">{{ log.bearbeiter }}</span>
            <div class="log-right">
              <span class="log-time">{{ new Date(log.timestamp).toLocaleString() }}</span>
              <OnyxButton
                label=""
                color="danger"
                mode="plain"
                :icon="iconTrash"
                @click="emit('removeLog', log.id)"
              />
            </div>
          </div>
          <p class="log-body">{{ log.notiz }}</p>
        </div>
        <div v-if="!selectedLaptop.it_logs?.length" class="empty-state">
          Bisher keine Support-Einträge für dieses Gerät hinterlegt.
        </div>
      </div>
      <!-- Neuanlage Log -->
      <form @submit.prevent="handleCreateLog" class="log-form">
        <OnyxHeadline is="h4">Neuen Support-Eintrag verfassen</OnyxHeadline>
        <OnyxInput label="Bearbeiter Name/ID" v-model="bearbeiter" required />
        <OnyxInput label="Durchgeführte Maßnahme / Update" v-model="notiz" required />
        <OnyxButton label="Eintrag Speichern" type="submit" color="primary" :icon="iconPlus" />
      </form>
    </div>
    <div v-else class="sd-card placeholder-card">
      <h3>Keine Auswahl getroffen</h3>
      <p>Wähle ein Gerät aus der linken Liste aus, um Details einzusehen und Logs zu verwalten.</p>
    </div>
  </section>
</template>

<style scoped>
.sd-card {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: var(--onyx-radius-md);
  padding: var(--onyx-spacing-lg);
}
.detail-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}
.device-id {
  font-size: 12px;
  font-weight: 700;
  color: #00a3e0;
}
.device-specs {
  display: flex;
  gap: 16px;
  margin-top: 6px;
  font-size: 14px;
  color: #4a5568;
}
.sd-divider {
  height: 1px;
  background-color: #edf2f7;
  margin: var(--onyx-spacing-lg) 0;
}
.logs-wrapper {
  display: flex;
  flex-direction: column;
  gap: var(--onyx-spacing-sm);
  margin: var(--onyx-spacing-md) 0;
}
.log-card {
  background: #f8fafc;
  border: 1px solid #e2e8f0;
  padding: var(--onyx-spacing-md);
  border-radius: var(--onyx-radius-sm);
}
.log-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
.log-author {
  font-weight: 700;
  font-size: 13px;
  color: #051829;
}
.log-right {
  display: flex;
  align-items: center;
  gap: 8px;
}
.log-time {
  font-size: 11px;
  color: #a0aec0;
}
.log-body {
  margin-top: 6px;
  font-size: 14px;
  color: #2d3748;
}
.log-form {
  display: flex;
  flex-direction: column;
  gap: var(--onyx-spacing-sm);
  margin-top: var(--onyx-spacing-xl);
  padding-top: var(--onyx-spacing-md);
  border-top: 1px dashed #cbd5e0;
}
.placeholder-card {
  text-align: center;
  padding: var(--onyx-spacing-2xl);
  color: #718096;
}
</style>