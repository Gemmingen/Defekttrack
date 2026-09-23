<script setup lang="ts">
import { onMounted } from 'vue';
import { useDefektTrack } from './composables/useDefektTrack';
import DigitsHeader from './components/DigitsHeader.vue';
import LaptopSidebar from './components/LaptopSidebar.vue';
import LaptopDetail from './components/LaptopDetail.vue';

const {
  laptops, selectedLaptop, activeFilter, isLoading, error,
  fetchLaptops, createLaptop, removeLaptop, addLog, removeLog
} = useDefektTrack();

const categories = ['Alle', 'Hardware', 'Display', 'Akku', 'Software', 'Netzwerk', 'Sonstiges'];

onMounted(() => fetchLaptops());

const handleFilterChange = (val: string) => {
  activeFilter.value = val;
  fetchLaptops(val);
};

const handleAddLog = (logData: { bearbeiter: string; notiz: string }) => {
  if (!selectedLaptop.value) return;
  addLog(selectedLaptop.value.id, logData);
};

const handleRemoveLog = (logId: number) => {
  if (!selectedLaptop.value) return;
  removeLog(selectedLaptop.value.id, logId);
};
</script>

<template>
  <div class="sd-app">
    <DigitsHeader />

    <main class="sd-main">
      <div v-if="error" class="error-banner">{{ error }}</div>

      <div class="dashboard-grid">
        <LaptopSidebar
          :laptops="laptops"
          :selected-laptop="selectedLaptop"
          :active-filter="activeFilter"
          :categories="categories"
          :is-loading="isLoading"
          @update:active-filter="handleFilterChange"
          @select-laptop="selectedLaptop = $event"
          @create-laptop="createLaptop"
        />

        <LaptopDetail
          :selected-laptop="selectedLaptop"
          @remove-laptop="removeLaptop"
          @add-log="handleAddLog"
          @remove-log="handleRemoveLog"
        />
      </div>
    </main>
  </div>
</template>

<style scoped>
:root {
  --onyx-font-family: 'Inter', system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
}

.sd-app {
  min-height: 100vh;
  background-color: #f4f6f8;
  font-family: 'Inter', system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  color: #1a202c;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

.sd-main {
  max-width: 1400px;
  margin: var(--onyx-spacing-xl) auto;
  padding: 0 var(--onyx-spacing-md);
}

.dashboard-grid {
  display: grid;
  grid-template-columns: 400px 1fr;
  gap: var(--onyx-spacing-xl);
}

.error-banner {
  background: #fff5f5;
  color: #c53030;
  border: 1px solid #feb2b2;
  padding: var(--onyx-spacing-md);
  border-radius: var(--onyx-radius-md);
  margin-bottom: var(--onyx-spacing-lg);
}
</style>