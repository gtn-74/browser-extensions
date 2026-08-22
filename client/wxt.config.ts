import { defineConfig } from 'wxt';
import { tanstackRouter } from '@tanstack/router-plugin/vite';

// See https://wxt.dev/api/config.html
export default defineConfig({
  modules: ['@wxt-dev/module-react'],
  vite: () => ({
    plugins: [
      tanstackRouter({
        target: 'react',
        routesDirectory: 'entrypoints/sidepanel/routes',
        generatedRouteTree: 'entrypoints/sidepanel/routeTree.gen.ts',
        autoCodeSplitting: false,
      }) as any,
    ],
  }),
  manifest: {
    name: 'WXT × Go Sidepanel Lab',
    description: 'ブラウザ右サイドのReact UIからGo APIを叩くサンプル',
    // 右サイドのサイドパネルを使うための権限
    permissions: ['sidePanel', 'storage'],
    // サイドパネルからローカルのGo APIへfetchするための権限（CORS回避にもなる）
    host_permissions: ['http://localhost:8080/*'],
  },
});
