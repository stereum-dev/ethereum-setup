<template>
  <div>
    <span> Setup SSH Connection to Server </span>
    <div class="text-left">
      <b-form-input v-model="model.host" placeholder="IP or Hostname"></b-form-input>
      <b-form-input v-model="model.user" placeholder="Username"></b-form-input>
      <b-form-checkbox v-model="model.sshKeyAuth" name="check-button" switch>
                Use SSH key authentication
      </b-form-checkbox>
      <b-form-input type="password" name="user" v-model="model.password" placeholder="Password" v-if="model.sshKeyAuth == false" />
      <b-form-input type="text" name="sshKeyAuthFile" v-model="model.keyfileLocation" placeholder="Keyfile Location" v-if="model.sshKeyAuth" />            
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
