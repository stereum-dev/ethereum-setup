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
        host: '',
        keyfileLocation: '',
        sshKeyAuth: false,
      }      
    };
  },
  components: {
    InstallationParameters,
  },
};
</script>
