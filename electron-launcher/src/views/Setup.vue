<template>
  <div class="home">    
    <installation-parameters :model="connectionParms" />
  </div>
</template>

<script>
// @ is an alias to /src
import InstallationParameters from "@/components/InstallationParameters.vue";

export default {
  name: "Setup",
  data() {
    return {
      connectionParms: {
        host: '65.21.178.198',
        keyfileLocation: 'C:\\Users\\Michael\\Documents\\develop\\rl\\electron-launcher\\src\\id_rsa',
      }      
    };
  },
  components: {
    InstallationParameters,
  },
};
</script>
