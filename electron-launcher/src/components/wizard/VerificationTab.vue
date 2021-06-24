<template>
  <div>
    <span v-if="!running">
      Please check the entered information and start the installation
    </span>
    <data>
      <div v-if="!running || done">
        <table class="text-left table">
          <tr>
            <td class="table-label">Host:</td>
            <td class="table-data">{{ model.host }}</td>
          </tr>
          <tr>
            <td class="table-label">User:</td>
            <td class="table-data">{{ model.user }}</td>
          </tr>
          <tr>
            <td class="table-label">Authentication:</td>
            <td v-if="model.password" class="table-data">Password</td>
            <td v-if="model.keyfileLocation" class="table-data">PEM Key from {{ model.keyfileLocation }}</td>
          </tr>
                    <tr>
            <td class="table-label">Stereum Release:</td>
            <td class="table-data">{{ model.stereumRelease }}</td>
          </tr>          
        </table>
      </div>

      <div v-if="running || done">
        <div v-if="running">
          <div class="alert alert-primary" role="alert">
            <b>Installation in Progress, please be patient</b>&nbsp;<i
              class="fas fa-cog fa-spin"
            ></i>            
          </div>
        </div>

        <div v-if="success === true">
          <div class="alert alert-success" role="alert">
            Installation Successful! You are good to go!
          </div>
        </div>
        <div v-if="success === false">
          <div class="alert alert-danger" role="alert">
            Unfortunately the Installation failed, please consult logs for
            details
          </div>
        </div>        
      </div>
    </data>
  </div>
</template>

<script>

export default {
  name: "VerificationTab",
  components: { },
  props: {
    model: Object,
    running: Boolean,
    success: Boolean,
    done: Boolean,
    progress: Number,
    logs: Array,
  },
  data() {
    return {
      max: 100,
    };
  },
  methods: {},
};
</script>

<style scoped>
.alert-primary {
  color: #fff;
  background-color: rgb(51, 102, 102);
  border-color: rgb(51, 102, 102);
}
</style>
