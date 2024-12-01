import { defineConfig, loadEnv } from "vite";
import vue from "@vitejs/plugin-vue";

// https://vitejs.dev/config/
export default defineConfig(({ command, mode }) => {
  const env = loadEnv(mode, process.cwd(), "");

  return {
    plugins: [vue()],
    server: {
      host: "0.0.0.0",
      hmr: {
        clientPort: env.OUTER_PORT_FRONTEND,
      },
      port: env.INNER_PORT_FRONTEND,
      // watch: {
      //   usePolling: true,
      //   interval: 1000,
      // }
    },
  };
});
