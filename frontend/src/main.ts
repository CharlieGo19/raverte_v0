import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import * as Wails from '@wailsapp/runtime';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.min.js';

Wails.Init(() => {
    createApp(App).use(router).mount('#app');
});
