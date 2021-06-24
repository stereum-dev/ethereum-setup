<template>
  <div class="home">    
    <!--<installation-parameters :model="connectionParms" />-->
    <setup-wizard :model="connectionParms"></setup-wizard>
  </div>
</template>

<script>
// @ is an alias to /src
import InstallationParameters from "@/components/InstallationParameters.vue";
import SetupWizard from "@/components/SetupWizard.vue";

export default {
  name: "Setup",
  components: {
    InstallationParameters,
    SetupWizard,
  },
  data() {
    return {
      connectionParms: {
        host: '',
        keyfileLocation: '',
        sshKeyAuth: false,
      }
    };
  },  
};
</script>
