<template>
  <div>
    <b-card
      img-src="setup-banner.png"
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
      stereumStatus: [],
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

      let status = 0;
      if (this.model.stereumRelease != this.stereumStatus.ccwebRelease) {
        //launch setup
        try {
              status = await ControlService.setup({ stereumRelease: this.model.stereumRelease });            
        } catch (ex) {
          console.log(ex);
          this.$toasted.show('Error setting up stereum launcher! Level: ' + ex.level + " Message: " + ex.message);
          return;
        }
      }

      if (status == 0) {
        await ControlService.openTunnels([{ dstPort: 8000, localPort: 8081}, { dstPort: 3000, localPort: 8082}, { dstPort: 7500, localPort: 8083}]);
        window.location.replace("http://localhost:8081/public/control-center");
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
        this.stereumStatus = await ControlService.inquire(this.model);
        this.releases = [];
        this.model.stereumRelease = undefined;
        if (!this.stereumStatus.exists) {
          this.releases = [this.stereumStatus.latestRelease]
        } else {
          if (this.stereumStatus.latestRelease) {
            this.releases.push({ text: `${this.stereumStatus.latestRelease} (latest)`, value: this.stereumStatus.latestRelease });
            this.model.stereumRelease = this.stereumStatus.latestRelease;
          }

          if (this.stereumStatus.latestRcRelease)
            this.releases.push({ text: `${this.stereumStatus.latestRcRelease} (latest unstable - not recommended!)`, value: this.stereumStatus.latestRcRelease });
          
          if (this.stereumStatus.existingRelease) {
            this.releases.push({ text: `${this.stereumStatus.existingRelease} (installed node)`, value: this.stereumStatus.existingRelease });
            this.model.stereumRelease = this.stereumStatus.existingRelease;
          }
          
          if (this.stereumStatus.ccRelease) {
            this.releases.push({ text: `${this.stereumStatus.ccRelease} (installed control-center)`, value: this.stereumStatus.ccRelease });
            if (this.model.stereumRelease === undefined) {
              this.model.stereumRelease = this.stereumStatus.ccRelease;
            }
          }

          if (this.stereumStatus.ccwebRelease) {
            this.releases.push({ text: `${this.stereumStatus.ccwebRelease} (running control-center)`, value: this.stereumStatus.ccwebRelease });
            if (this.model.stereumRelease === undefined) {
              this.model.stereumRelease = this.stereumStatus.ccwebRelease;
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
