<template>
  <div>
    <h3>Choose Version - multiple found</h3>

    <div class="text-left">
      <b-form-radio-group
        :stacked="true"
        label="Pick Stereum Release"
        :options="releases"
        v-model="model.stereumRelease"
        value-field="value"
      >
      </b-form-radio-group>      
    </div>
  </div>
</template>

<script>
export default {
  name: "ReleaseTab",
  components: {},  
  props: {
    releases: Array,
    model: Object,
  },
};
</script>

<style scoped></style>
