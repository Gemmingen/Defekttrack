<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { OnyxHeadline, OnyxButton, OnyxSelect, OnyxInput } from 'sit-onyx';
import { iconTrash, iconPlus } from '@sit-onyx/icons';
import { useDefektTrack } from './composables/useDefektTrack';
import LaptopForm from './components/LaptopForm.vue';

const { 
  laptops, selectedLaptop, activeFilter, isLoading, error, 
  fetchLaptops, createLaptop, removeLaptop, addLog, removeLog 
} = useDefektTrack();

const bearbeiter = ref('');
const notiz = ref('');
const categories = ['Alle', 'Hardware', 'Display', 'Akku', 'Software', 'Netzwerk', 'Sonstiges'];

onMounted(() => fetchLaptops());

const handleFilterChange = (val: string) => {
  activeFilter.value = val;
  fetchLaptops(val);
};

const handleCreateLog = () => {
  if (!selectedLaptop.value || !bearbeiter.value || !notiz.value) return;
  addLog(selectedLaptop.value.id, { bearbeiter: bearbeiter.value, notiz: notiz.value });
  notiz.value = '';
};
</script>

<template>
  <div class="sd-app">
    <!-- Digits Header -->
    <header class="sd-navbar">
      <div class="nav-container">
        <div class="brand">
          <span class="brand-accent">SCHWARZ DIGITS</span>
          <OnyxHeadline is="h1" class="brand-title">DefektTrack</OnyxHeadline>
        </div>
        <div class="nav-status">
          <span class="status-indicator"></span>
          <span>Enterprise Support Portal</span>
        </div>
      </div>
    </header>

    <main class="sd-main">
      <div v-if="error" class="error-banner">{{ error }}</div>

      <div class="dashboard-grid">
        <!-- Linke Spalte: Erfassung & Filter-Liste -->
        <aside class="sidebar-panel">
          <LaptopForm @submit="createLaptop" />

          <div class="sd-card filter-card">
            <OnyxSelect 
              label="Filter nach Kategorie" 
              listLabel="Kategorien auswählen"
              v-model="activeFilter" 
              :options="categories.map(c => ({ label: c, value: c }))"
              @update:model-value="handleFilterChange"
            />
          </div>

          <div class="list-container">
            <div v-if="isLoading" class="loading-text">Lade Datensätze...</div>
            <div v-else class="laptop-list">
              <div 
                v-for="laptop in laptops" 
                :key="laptop.id" 
                class="laptop-card"
                :class="{ active: selectedLaptop?.id === laptop.id }"
                @click="selectedLaptop = laptop"
              >
                <div class="card-head">
                  <strong>{{ laptop.marke }} {{ laptop.name }}</strong>
                  <span class="tag-category">{{ laptop.fehler }}</span>
                </div>
                <div class="card-sub">{{ laptop.os }}</div>
                <div class="card-foot">
                  <span>{{ laptop.it_logs?.length || 0 }} Support-Einträge</span>
                </div>
              </div>
            </div>
          </div>
        </aside>

        <!-- Rechte Spalte: Laptop Details und Support-Logs -->
        <section class="detail-panel">
          <div v-if="selectedLaptop" class="sd-card detail-card">
            <div class="detail-header">
              <div>
                <span class="device-id">ID #{{ selectedLaptop.id }}</span>
                <OnyxHeadline is="h2">{{ selectedLaptop.marke }} {{ selectedLaptop.name }}</OnyxHeadline>
                <div class="device-specs">
                  <span><strong>Kategorie:</strong> {{ selectedLaptop.fehler }}</span>
                  <span><strong>OS:</strong> {{ selectedLaptop.os }}</span>
                </div>
              </div>
              <OnyxButton 
                label="Löschen" 
                color="danger" 
                mode="outline" 
                :icon="iconTrash" 
                @click="removeLaptop(selectedLaptop.id)" 
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
                      @click="removeLog(selectedLaptop.id, log.id)" 
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
/* Schwarz Digits Navbar */
.sd-navbar {
  background-color: #051829;
  border-bottom: 3px solid #00a3e0;
  padding: var(--onyx-spacing-md) var(--onyx-spacing-2xl);
  color: white;
}

.nav-container {
  max-width: 1400px;
  margin: 0 auto;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.brand {
  display: flex;
  flex-direction: column;
}

.brand-accent {
  font-size: 10px;
  font-weight: 800;
  letter-spacing: 2px;
  color: #00e5ff;
}

.brand-title {
  color: #ffffff;
  margin: 0;
  font-size: 24px;
}

.nav-status {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: #a0aec0;
}

.status-indicator {
  width: 8px;
  height: 8px;
  background-color: #00e5ff;
  border-radius: 50%;
  box-shadow: 0 0 8px #00e5ff;
}

/* Layout */
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

.sidebar-panel {
  display: flex;
  flex-direction: column;
  gap: var(--onyx-spacing-md);
}

.sd-card {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: var(--onyx-radius-md);
  padding: var(--onyx-spacing-lg);
}

.laptop-list {
  display: flex;
  flex-direction: column;
  gap: var(--onyx-spacing-xs);
}

.laptop-card {
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: var(--onyx-radius-md);
  padding: var(--onyx-spacing-md);
  cursor: pointer;
  transition: all 0.2s ease;
}

.laptop-card:hover {
  border-color: #00a3e0;
}

.laptop-card.active {
  border-color: #051829;
  border-left: 4px solid #00a3e0;
  background: #fafafa;
}

.card-head {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.tag-category {
  background: #e6f7ff;
  color: #005b9a;
  font-size: 11px;
  font-weight: 700;
  padding: 2px 8px;
  border-radius: 4px;
}

.card-sub {
  font-size: 13px;
  color: #718096;
  margin-top: 2px;
}

.card-foot {
  font-size: 11px;
  color: #a0aec0;
  margin-top: 8px;
}

/* Detailbereich */
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

.error-banner {
  background: #fff5f5;
  color: #c53030;
  border: 1px solid #feb2b2;
  padding: var(--onyx-spacing-md);
  border-radius: var(--onyx-radius-md);
  margin-bottom: var(--onyx-spacing-lg);
}
</style>