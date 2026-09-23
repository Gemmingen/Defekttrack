import { createApp } from 'vue'
import App from './App.vue'

import 'sit-onyx/style.css'

const fontStyle = document.createElement('style')
fontStyle.innerHTML = `
  @import url('https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700;800&display=swap');

  :root {
    --onyx-font-family: 'Inter', sans-serif !important;
    --onyx-font-family-headline: 'Inter', sans-serif !important;
    --onyx-font-family-body: 'Inter', sans-serif !important;
  }

  *, *::before, *::after,
  h1, h2, h3, h4, h5, h6, p, span, div, button, input, select, textarea,
  [class*="onyx-"] {
    font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif !important;
  }
`
document.head.appendChild(fontStyle)

createApp(App).mount('#app')