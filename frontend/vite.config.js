import { defineConfig } from 'vite';

export default defineConfig({
    server: {
        port: 5173, // Port untuk frontend Anda
        proxy: {
            '/api': {
                target: 'http://localhost:8080', // Target backend Anda
                changeOrigin: true,
                secure: false,
            },
        },
    },
});
