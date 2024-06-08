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
        protocol: "ws",
        port: env.HMR_WS_PORT
      },
      port: env.INNER_PORT_FRONTEND,
    },
  };
});
