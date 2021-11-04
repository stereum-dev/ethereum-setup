<template>
  <div class="container text-left">
    <form @submi.prevent="submitForm" autocomplete="off">
      <h3>Connecting to your server</h3>

      <div class="row">
        <div class="col-9 pt-2 pb-2">
          <b-form-select id="saved-connections" v-model="selectedConnection" :options="savedConnections"></b-form-select>
          <b-tooltip target="saved-connections" triggers="hover">Saved connections</b-tooltip>
        </div>
        <div class="col-3 pt-2 pb-2">
          <template v-if="showAdd">
            <b-button
              id="add-connection"
              size="sm"
              @click="addItem"
              variant="primary"
              :disabled="validValidation"
              class="mb-2 mr-sm-2 mb-sm-0 mt-1"
            >
              <b-icon icon="plus" aria-hidden="true" title=""></b-icon>
            </b-button>
            <b-tooltip target="add-connection" triggers="hover">Add currenct connection as new connection (password won't get stored)</b-tooltip>
          </template>
          <template v-if="showEdit">
            <b-button
              id="edit-connection"
              size="sm"
              @click="editItem"
              class="mb-2 mr-sm-2 mb-sm-0 mt-1"
            >
              <b-icon icon="gear-fill" aria-hidden="true" title=""></b-icon>
            </b-button>
            <b-tooltip target="edit-connection" triggers="hover">Edit currenct connection</b-tooltip>
          </template>
          <template v-if="showSave">
            <b-button
              id="save-connection"
              size="sm"
              @click="saveItem"
              variant="success"
              :disabled="validValidation"
              class="mb-2 mr-sm-2 mb-sm-0 mt-1"
            >
              <b-icon icon="check" aria-hidden="true" title=""></b-icon>
            </b-button>
            <b-tooltip target="save-connection" triggers="hover">Save changes</b-tooltip>
          </template>
          <template v-if="showRemove">
            <b-button
              id="remove-connection"
              size="sm"
              @click="removeItem"
              variant="danger"
              class="mb-2 mr-sm-2 mb-sm-0 mt-1"
            >
              <b-icon icon="trash" aria-hidden="true"></b-icon>
            </b-button>
            <b-tooltip target="remove-connection" triggers="hover">Delete selected connection</b-tooltip>
          </template>
        </div>
      </div>

      <template v-if="showDetails">
        <div class="row">
          <div class="col-12 pt-2 pb-2">
            <b-form-input
              id="servername" 
              v-model="model.servername" 
              :state="servernameValid" 
              placeholder="Give your server a name"
            ></b-form-input>
            <b-tooltip target="servername" triggers="hover">This server name will be displayed in the dropdown menu</b-tooltip>
          </div>
        </div>
        <div class="row">
          <div class="col-10 pt-2 pb-2">
            <b-form-input 
              id="host" 
              v-model="model.host" 
              :state="hostValid" 
              placeholder="IP or Hostname"
            ></b-form-input>
            <b-tooltip target="host" triggers="hover">IP or Hostname</b-tooltip>
          </div>
          <div class="col-2 pt-2 pb-2">
            <b-form-input 
              id="port" 
              v-model="model.port" 
              :state="portValid" 
              placeholder="Port (default 22)"
            ></b-form-input>
            <b-tooltip target="port" triggers="hover">SSH Port (default 22)</b-tooltip>
          </div>
        </div>
        <div class="row">
          <div class="col-12 pt-2 pb-2">
            <b-form-input 
              id="user" 
              v-model="model.user" 
              :state="userValid" 
              placeholder="Username"
            ></b-form-input>
            <b-tooltip target="user" triggers="hover">Username</b-tooltip>
          </div>
        </div>
      </template>

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
    </form>
  </div>
</template>

<script>
import ControlService from "@/store/ControlService";

export default {
  name: "ConnectionTab",
  components: {

  },
  data() {
    const unselectedConnection = {
          value: null,
          text: "add server",
        };

    return {
      selectedConnection: null,
      unselectedConnection: unselectedConnection,
      savedConnections: [unselectedConnection],
      selectionInProgress: false,
      showDetails: true,
      showAdd: true,
      showEdit: false,
      showRemove: false,
      showSave: false,
    };
  },
  props: {
    model: Object, 
  },
  watch: {
    "selectedConnection": function() {
      if (!this.selectionInProgress && this.selectedConnection !== this.unselectedConnection.value) {
        this.selectionInProgress = true;

        this.model.servername = this.selectedConnection.servername;
        this.model.host = this.selectedConnection.host;
        this.model.user = this.selectedConnection.user;
        this.model.keyfileLocation = this.selectedConnection.keyfileLocation;
        this.model.sshKeyAuth = this.selectedConnection.sshKeyAuth;
        this.model.port = this.selectedConnection.port;
        this.showDetails = false;
        this.showAdd = false;
        this.showEdit = true;
        this.showRemove = true;
        this.showSave = false;

        this.selectionInProgress = false;
      } else {
        this.model.servername = "";
        this.model.host = "";
        this.model.user = "";
        this.model.keyfileLocation = "";
        this.model.sshKeyAuth = false;
        this.model.port = 22;
        this.showDetails = true; 
        this.showAdd = true;
        this.showEdit = false;
        this.showRemove = false;
        this.showSave = false;
      }
    },
  },

  computed: {
    servernameValid() { 
      if (this.model.servername.length > 0){
        return true;
      } else {
        return false;
      }
    },
    hostValid() {
      if (this.model.host.length > 0){
        return true;
      } else {
        return false;
      }
    },
    portValid() {
      if (this.model.port.length < 1){
        return false;
      } else {
        return true;
      }
    },
    userValid() {
      if (this.model.user.length > 0){
        return true;
      } else {
        return false;
      }
    },
    validValidation() {
      return !(this.servernameValid && this.hostValid && this.portValid && this.userValid)
    },
  },
  
  created() {
    this.loadStoredConnections();
  },

  methods: {
    removeItem() {

      if (this.selectedConnection !== this.unselectedConnection.value) {
        this.selectionInProgress = true;

        const selectedConn = this.selectedConnection;

        this.savedConnections = this.savedConnections.filter(function(conn) {
          return selectedConn != conn.value;
        });

        this.selectedConnection = this.unselectedConnection.value;

        this.selectionInProgress = false;

        ControlService.writeConfig({ savedConnections: this.storableConnections() });
      }

      this.model.servername = "";
      this.model.host = "";
      this.model.user = "";
      this.model.keyfileLocation = "";
      this.model.sshKeyAuth = false;
      this.model.port = 22;
    },
    addItem() {
      this.selectionInProgress = true;

      const newConn = this.createConnectionModel(this.model.servername, this.model.host, this.model.port, this.model.user, this.model.sshKeyAuth, this.model.keyfileLocation);

      this.savedConnections.push({
        value: newConn,
        text: newConn.servername ? newConn.servername : newConn.host + " (" + newConn.user + ")",
      });

      this.selectedConnection = newConn;

      this.selectionInProgress = false;

      ControlService.writeConfig({ savedConnections: this.storableConnections() });
    },
    editItem() {
      this.showDetails = !this.showDetails;
      this.showAdd = false;
      this.showEdit = false;
      this.showRemove = true;
      this.showSave = true;
    },
    saveItem() {
      if (this.selectedConnection !== this.unselectedConnection.value) {
        this.selectionInProgress = true;

        this.selectedConnection.servername = this.model.servername;
        this.selectedConnection.host = this.model.host;
        this.selectedConnection.user = this.model.user
        this.selectedConnection.keyfileLocation = this.model.keyfileLocation;
        this.selectedConnection.sshKeyAuth = this.model.sshKeyAuth;
        this.selectedConnection.port = this.model.port;

        this.selectionInProgress = false;
        ControlService.writeConfig({ savedConnections: this.storableConnections() });
        this.loadStoredConnections();
        this.showSave = false;
        this.showDetails = false;
        this.showEdit = true;
      }
    },
    createConnectionModel(servername, host, port, user, sshKeyAuth, keyfileLocation) {
      return {
        servername: servername,
        host: host,
        port: port,
        user: user,
        sshKeyAuth: sshKeyAuth,
        keyfileLocation: keyfileLocation,
      };
    },
    storableConnections() {
      return this.savedConnections.filter(function(conn) {
          return conn.value != null && conn.value.value === undefined;
      }).map(conn2=>{
        return conn2.value;
      }); 
    },
    loadStoredConnections: async function() {
      const storageSavedConnections = await ControlService.readConfig();

      //let savedConnections = [this.unselectedConnection];
      let savedConnections = [];

      if (storageSavedConnections !== undefined && storageSavedConnections.savedConnections !== undefined) {
        savedConnections = savedConnections.concat(storageSavedConnections.savedConnections);
      }

      this.savedConnections = [this.unselectedConnection].concat(savedConnections.map((conn2, index)=>{
        return {
          text: conn2.servername ? conn2.servername : /*conn2.host + " (" + conn2.user + ")"*/ "Server " + (index + 1), 
          value: conn2,
        }
      }));
    },
  }
};
</script>


<style scoped></style>