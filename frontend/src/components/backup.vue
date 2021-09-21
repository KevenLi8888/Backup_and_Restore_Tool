<template>
  <v-container fluid class="px-0">
    <v-row>
      <v-col cols="12" sm="6" offset-sm="3">
        <v-card outlined class="pa-4 ma-4">
          <v-card-title>
            Select source
          </v-card-title>
          <v-card-text>
            {{srcPath}}
          </v-card-text>
          <v-card-actions>
            <v-layout justify-center align-center class="px-0">
              <v-btn color="#81D4FA" @click="selectSrcDir">Open...</v-btn>
            </v-layout>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12" sm="6" offset-sm="3">
        <v-card outlined class="pa-4 ma-4">
          <v-card-title>
            Advanced Features
          </v-card-title>
          <v-card-actions>
            <v-layout justify-center align-center class="px-0">
              <v-btn color="blue" @click="getMessage">Press Me!!!</v-btn>
            </v-layout>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12" sm="6" offset-sm="3">
        <v-card outlined class="pa-4 ma-4">
          <v-card-title>
            Select destination
          </v-card-title>
          <v-card-text>
            {{destPath}}
          </v-card-text>
          <v-card-actions>
            <v-layout justify-center align-center class="px-0">
              <v-btn color="#81D4FA" @click="selectDestDir">Open...</v-btn>
            </v-layout>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12" sm="6" offset-sm="3">
        <v-layout justify-center align-center class="px-0">
          <v-btn fab color="success" disabled="!activateButton">
            <v-icon>
              backup
            </v-icon>
          </v-btn>
        </v-layout>
      </v-col>
    </v-row>

    <div class='text-xs-center'>
      <v-dialog v-model="errorDialog" width="500">
        <v-card>
          <v-card-title class="headline" primary-title>ERROR!</v-card-title>
          <v-card-text>{{errorMessage}}</v-card-text>
          <v-divider></v-divider>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" text @click="errorDialog = false">Fine</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
    <div class="text-xs-center">
      <v-dialog v-model="dialog" width="500">
        <v-card>
          <v-card-title class="headline" primary-title>Message from Go</v-card-title>
          <v-card-text>{{message}}</v-card-text>
          <v-divider></v-divider>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" text @click="dialog = false">Awesome</v-btn>
<!--            TODO:待优化，只保留一个考虑message是否要清空 -->
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </v-container>
</template>

<script>
  export default {
    data () {
      return {
        message: "",
        raised: true,
        dialog: false,
        errorDialog: false,
        errorMessage: " ",
        srcPath: "",
        destPath: "",
        activateButton: false
      }
    },

    computed: {
      listenChange() {
        const srcPath = this.srcPath
        const destPath = this.destPath
        return {srcPath, destPath}
      }
    },

    watch: {
      listenChange:{
        handler:function(val) {
          console.log('listening change' + val.srcPath + val.destPath)
          var self = this
          if (val.srcPath != "" && val.destPath != "") {
            self.activateButton = true
        }
      },
        deep: true
      }
    },

    methods: {
      selectSrcDir: function () {
        var self = this
        window.backend.Backup.SelectSourceDir()
            .then(result => {
              var self = this
              self.srcPath = result
            })
            .catch(error => {
              if (self.srcPath == "") {
                self.errorMessage = error
                self.errorDialog = true
                setTimeout(() => {
                  self.errorMessage = ""
                  self.errorDialog = false
                }, 5000)
              }
            });
      },

      selectDestDir: function () {
        var self = this
        window.backend.Backup.SelectDestDir()
            .then(result => {
              var self = this
              self.destPath = result
            })
            .catch(error => {
              if (self.destPath == "") {
                self.errorMessage = error
                self.errorDialog = true
                setTimeout(() => {
                  self.errorMessage = ""
                  self.errorDialog = false
                }, 5000)
              }
            });
      },

      getMessage: function () {
        var self = this
        window.backend.basic().then(result => {
          self.message = result
          self.dialog = true
        })
      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  h1 {
    margin-top: 2em;
    position: relative;
    min-height: 5rem;
    width: 100%;
  }

  a:hover {
    font-size: 1.7em;
    border-color: blue;
    background-color: blue;
    color: white;
    border: 3px solid white;
    border-radius: 10px;
    padding: 9px;
    cursor: pointer;
    transition: 500ms;
  }

  a {
    font-size: 1.7em;
    border-color: white;
    background-color: #121212;
    color: white;
    border: 3px solid white;
    border-radius: 10px;
    padding: 9px;
    cursor: pointer;
  }


</style>
