<template>
  <div class="container text-left">
    <h3>Connecting to your server</h3>

    <div class="row">
        <div class="col-12 pt-2 pb-2">
          <b-form-input v-model="model.host" placeholder="IP or Hostname"></b-form-input>
        </div>
    </div>
    <div class="row">
        <div class="col-12 pt-2 pb-2">
          <b-form-input v-model="model.port" placeholder="Port (default 22)"></b-form-input>
        </div>
    </div>
    <div class="row">
        <div class="col-12 pt-2 pb-2">
          <b-form-input v-model="model.user" placeholder="Username"></b-form-input>
        </div>
    </div>
    <div class="row">
        <div class="col-12 pt-2 pb-2">
          <b-form-checkbox v-model="model.sshKeyAuth" name="check-button" switch>
                Use SSH key authentication
          </b-form-checkbox>
        </div>
    </div>
    <div class="row" v-if="model.sshKeyAuth == false">
        <div class="col-12 pt-2 pb-2">
          <b-form-input type="password" name="user" v-model="model.password" placeholder="Password" />
        </div>
    </div>
    <div class="row" v-if="model.sshKeyAuth">
        <div class="col-12 pt-2 pb-2">
          <b-form-input type="text" name="sshKeyAuthFile" v-model="model.keyfileLocation" placeholder="Keyfile Location" />
        </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ConnectionTab",
  components: {},  
  props: {
    model: Object,
  },
};
</script>

<style scoped></style>
