<script setup lang="ts">
import { OnyxHeadline } from 'sit-onyx';
</script>

<template>
  <header class="sd-navbar">
    <div class="nav-container">
      <div class="brand">
        <span class="brand-accent">SCHWARZ DIGITS</span>
        <OnyxHeadline is="h1" class="brand-title">DefektTrack</OnyxHeadline>
      </div>
      <div class="nav-status">
        <span class="status-indicator"></span>
        <span>Service Point</span>
      </div>
    </div>
  </header>
</template>

<style scoped>
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
</style>