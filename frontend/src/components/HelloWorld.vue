<script setup lang="ts">
import { ref } from 'vue'
import { PingUnaryRequest } from '../../proto-gen-web/backend/messages_pb';
import { HealthcheckServiceClient } from '../../proto-gen-web/backend/ServicesServiceClientPb';
defineProps<{ msg: string }>()

const count = ref(0)
// it is recommend to use one client between modules.
const healthcheckClient = new HealthcheckServiceClient(
  "http://localhost:9000" // request to envoy proxy
);
const req = new PingUnaryRequest();
req.setPing("ping");
const metadata = {}; // gRPC metadata(e.g. set jwt token)
healthcheckClient.pingUnary(req, metadata, (err, resp) => {
  if (err) {
    console.error(err);
  } else {
    console.log(resp.getPong());
  }
});
</script>

<template>
  <h1>{{ msg }}</h1>

  <div>
    <button class="border-black border-[1px] p-2" type="button" @click="count++">count is {{ count }}</button>
    <p>
      Edit
      <code>components/HelloWorld.vue</code> to test HMR
    </p>
  </div>

  <p>
    Check out
    <a href="https://vuejs.org/guide/quick-start.html#local" target="_blank">create-vue</a>, the official Vue + Vite
    starter
  </p>
  <p>
    Install
    <a href="https://github.com/vuejs/language-tools" target="_blank">Volar</a>
    in your IDE for a better DX
  </p>
  <p>Click on the Vite and Vue logos to learn more</p>
</template>
