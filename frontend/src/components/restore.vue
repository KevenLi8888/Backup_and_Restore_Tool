<template>
  <v-container fluid class="px-0">
    <v-row>
      <v-col cols="12" sm="6" offset-sm="3">
        <v-card outlined class="pa-4 ma-4">
          <v-card-title>
            选择需要还原的文件
          </v-card-title>
          <v-card-text>
            {{srcPath}}
          </v-card-text>
          <v-card-actions>
            <v-layout justify-center align-center class="px-0">
              <v-btn color="#81D4FA" @click="selectSrcFile">Open...</v-btn>
            </v-layout>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12" sm="6" offset-sm="3">
        <v-card outlined class="pa-4 ma-4">
          <v-card-title>
            输入密码
          </v-card-title>
          <v-text-field
              v-model="password"
              :append-icon="showPassowrd ? 'visibility' : 'visibility_off'"
              :rules="[rules.required, rules.max]"
              :type="showPassowrd ? 'text' : 'password'"
              name="input-10-1"
              label="输入密码"
              hint="密码不得多于15位"
              counter
              @click:append="showPassowrd = !showPassowrd"
          ></v-text-field>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12" sm="6" offset-sm="3">
        <v-card outlined class="pa-4 ma-4">
          <v-card-title>
            选择目标备份路径（未实现）
          </v-card-title>
          <v-card-text>
            {{destPath}}
          </v-card-text>
          <v-card-actions>
            <v-layout justify-center align-center class="px-0">
              <v-btn disabled color="#81D4FA" @click="selectDestDir">Open...</v-btn>
            </v-layout>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12" sm="6" offset-sm="3">
        <v-layout justify-center align-center class="px-0">
          <v-btn fab color="success" @click="confirmRestore">
            <v-icon>
              restore
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
  import Wails from "@wailsapp/runtime";

  export default {
    data () {
      return {
        message: "",
        raised: true,
        dialog: false,
        errorDialog: false,
        errorMessage: "",
        srcPath: "",
        destPath: "notUsed",
        // activateButton: false,
        showPassword: false,
        password:"",
        rules: {
          required: value => !!value || 'Required.',
          max: v => v.length < 16 || '至多输入15位密码'
        },
      }
    },


    methods: {
      selectSrcFile: function () {
        var self = this
        window.backend.Backup.SelectRestoreFile()
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

      confirmRestore: function () {
        var self = this
        if (self.srcPath == "") {
          self.errorMessage = "请选择需要还原的文件！"
          self.errorDialog = true
          setTimeout(() => {
            self.errorMessage = ""
            self.errorDialog = false
          }, 5000)
        } else if (self.password == "" || self.password.length > 15) {
          self.errorMessage = "请正确输入密码！"
          self.errorDialog = true
          setTimeout(() => {
            self.errorMessage = ""
            self.errorDialog = false
          }, 5000)
        } else {
          window.backend.Backup.PerformRestore(self.srcPath, self.password).then(result => {
            self.message = result
            self.dialog = true
          })
          .catch(error => {
            self.errorMessage = error
            self.errorDialog = true
            setTimeout(() => {
              self.errorMessage = ""
              self.errorDialog = false
            }, 5000)
          });
        }
      },

      getMessage: function () {
        var self = this
        window.backend.basic().then(result => {
          self.message = result
          self.dialog = true
        })
      }
    },
    mounted() {
      window.backend.Backup.LoadList().then((list) => {
        Wails.Log.Info("I got this list: " + list)
      });
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
