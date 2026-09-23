<script setup lang="ts">
import { OnyxSelect } from 'sit-onyx';
import LaptopForm from './LaptopForm.vue';
import {
  FILTER_OPTIONS,
  CATEGORY_LABELS,
  type FilterCategory,
} from '../constants/categories';
import type { Laptop, LaptopInput } from '../api/generated';

defineProps<{
  laptops: Laptop[];
  selectedLaptop: Laptop | null;
  activeFilter: FilterCategory;
  isLoading: boolean;
}>();

const emit = defineEmits<{
  (e: 'update:activeFilter', val: FilterCategory): void;
  (e: 'selectLaptop', laptop: Laptop): void;
  (e: 'createLaptop', data: LaptopInput): void;
}>();
</script>

<template>
  <aside class="sidebar-panel">
    <LaptopForm @submit="emit('createLaptop', $event)" />
    <div class="sd-card filter-card">
      <OnyxSelect
        label="Filter nach Kategorie"
        listLabel="Kategorien auswählen"
        :model-value="activeFilter"
        :options="FILTER_OPTIONS"
        @update:model-value="emit('update:activeFilter', $event as FilterCategory)"
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
          @click="emit('selectLaptop', laptop)"
        >
          <div class="card-head">
            <strong>{{ laptop.marke }} {{ laptop.name }}</strong>
            <span class="tag-category">{{ CATEGORY_LABELS[laptop.fehler] || laptop.fehler }}</span>
          </div>
          <div class="card-sub">{{ laptop.os }}</div>
          <div class="card-foot">
            <span>{{ laptop.it_logs?.length || 0 }} Support-Einträge</span>
          </div>
        </div>
      </div>
    </div>
  </aside>
</template>

<style scoped>
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
</style>