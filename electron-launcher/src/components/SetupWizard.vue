<template>
  <div>
    <b-card
      img-src="https://picsum.photos/600/50/?grayscale"
      img-alt="Image"
      img-top
      tag="article"
      style="max-width: 20rem"
      class="mb-2 main-card"
    >
      <b-card-body>
        <div>
          <form-wizard
            shape="square"
            color="#336666"
            title="Stereum Ethereum Node Setup Launcher"
            subtitle="Connect to your node"
            finishButtonText="Go!"
            @on-complete="onComplete"
            :hide-buttons="installationRunning || installationSuccess"
          >
            <tab-content title="Welcome" icon="faw fas fa-compass">
              <welcome-tab></welcome-tab>
            </tab-content>
            <tab-content title="Connection" icon="faw fas fa-network-wired" :before-change="() => fetchReleases()">
              <connection-tab :model="model"></connection-tab>
            </tab-content>
            <tab-content title="Release" icon="faw fas fa-cogs">
              <release-tab :releases="releases" :model="model"></release-tab>
            </tab-content>                        
            <tab-content title="Verify" icon="faw fas fa-check-double">
              <verification-tab
                :logs="logs.tasks"
                :model="model"
                :progress="installationProgress"
                :running="installationRunning"
                :success="installationSuccess"
                :done="installationDone"
              ></verification-tab>
            </tab-content>
          </form-wizard>
        </div>
      </b-card-body>
    </b-card>
  </div>
</template>

<script>
//local registration
import { FormWizard, TabContent } from "vue-form-wizard";
import "vue-form-wizard/dist/vue-form-wizard.min.css";
import "@fortawesome/fontawesome-free/css/all.css";
import WelcomeTab from "@/components/wizard/WelcomeTab.vue";
import ConnectionTab from "@/components/wizard/ConnectionTab.vue";
import ReleaseTab from "@/components/wizard/ReleaseTab.vue";
import VerificationTab from "@/components/wizard/VerificationTab.vue";
import ControlService from "@/store/ControlService";

export default {
  name: "SetupWizard",
  components: {
    FormWizard,
    TabContent,
    WelcomeTab,
    ConnectionTab,
    ReleaseTab,
    VerificationTab,    
  },
  data() {
    return {
      response: "",
      logs: {
        tasks: [],
      },
      installationRunning: false,
      installationDone: false,
      installationProgress: 0,
      installationSuccess: undefined,
      releases: [],
    };
  },
  props: {
    model: Object,
  },
  methods: {    
    onComplete: async function () {
      this.installationRunning = true;
      this.installationProgress = 0;      

      // write variables for the ansible call
      const extraVars = {        
        stereum_version_tag: window.STEREUM_VERSION_TAG,
      };      
      this.installationDone = false;          

      //launch setup
      try {
            // TODO: check if installation is necessary
            await ControlService.setup({ stereumRelease: this.model.stereumRelease });            
            // TODO: initiate tunnels if installation succeeded
      } catch (ex) {
         console.log(ex);
         this.$toasted.show('Error setting up stereum launcher! Level: ' + ex.level + " Message: " + ex.message);
         return;
      }
    },
    fetchReleases: async function () {
      try {
            await ControlService.connect(this.model);
      } catch (ex) {  
         console.log(ex);
         this.$toasted.show('Error connecting to server! Level: ' + ex.level + " Message: " + ex.message);
         return;
      }
      try {
        const stereumStatus = await ControlService.inquire(this.model);
        this.releases = [];
        this.model.stereumRelease = undefined;
        if (!stereumStatus.exists) {
          this.releases = [stereumStatus.latestRelease]
        } else {
          if (stereumStatus.latestRelease) {
            this.releases.push({ text: `${stereumStatus.latestRelease} (latest)`, value: stereumStatus.latestRelease });
            this.model.stereumRelease = stereumStatus.latestRelease;
          }

          if (stereumStatus.latestRcRelease)
            this.releases.push({ text: `${stereumStatus.latestRcRelease} (latest unstable - not recommended!)`, value: stereumStatus.latestRcRelease });
          
          if (stereumStatus.existingRelease) {
            this.releases.push({ text: `${stereumStatus.existingRelease} (installed node)`, value: stereumStatus.existingRelease });
            this.model.stereumRelease = stereumStatus.existingRelease;
          }
          
          if (stereumStatus.ccRelease) {
            this.releases.push({ text: `${stereumStatus.ccRelease} (installed control-center)`, value: stereumStatus.ccRelease });
            if (this.model.stereumRelease === undefined) {
              this.model.stereumRelease = stereumStatus.ccRelease;
            }
          }

          if (stereumStatus.ccwebRelease) {
            this.releases.push({ text: `${stereumStatus.ccwebRelease} (running control-center)`, value: stereumStatus.ccwebRelease });
            if (this.model.stereumRelease === undefined) {
              this.model.stereumRelease = stereumStatus.ccwebRelease;
            }
          }
        }
        return Promise.resolve(true);
      } catch (ex) {
         console.log(ex);
         this.$toasted.show('Error fetching Stereum Versions! Level: ' + ex.level + " Message: " + ex.message);
         return;
      }
    }
  },
};
</script>

<style scoped>
.main-card {
  margin-top: 5%;
  margin-left: 10%;
  margin-right: 10%;
  min-width: 80%;
}

.main-card img {
  max-height: 50px;
}
</style>
