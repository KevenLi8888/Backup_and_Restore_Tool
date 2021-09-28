<template>
  <v-container fluid class="px-0">

    <v-row>
      <v-col cols="12" sm="6" offset-sm="3">
        <v-card outlined class="pa-4 ma-4" height="85%">
          <v-card-title class="justify-center">
            选择需要还原的文件
          </v-card-title>
          <v-card-text class=text-center>
            {{srcPath}}
          </v-card-text>
          <v-card-actions>
            <v-layout justify-center align-center class="px-0">
              <v-btn color="#81D4FA" @click="selectSrcFile">打开...</v-btn>
            </v-layout>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>

    <v-row>
      <v-col cols="12" sm="6" offset-sm="3">
        <v-card outlined class="pa-4 ma-4" height="85%">
          <v-card-title class="justify-center">
            输入密码
          </v-card-title>
          <v-layout justify-center align-center class="px-0">
            <v-text-field
                v-model="password"
                :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
                :rules="[rules.required, rules.max]"
                :type="showPassword ? 'text' : 'password'"
                name="input-10-1"
                label="输入密码"
                outlined
                hint="密码不得多于15位"
                counter
                @click:append="showPassword = !showPassword"
            ></v-text-field>
          </v-layout>
        </v-card>
      </v-col>
    </v-row>

    <v-row>
      <v-col>
        <v-spacer></v-spacer>
      </v-col>
    </v-row>

    <v-row>
      <v-col cols="12" sm="6" offset-sm="3">
        <v-layout justify-center align-center class="px-0">
          <v-btn
              fab
              color="success"
              :loading="loading"
              :disabled="loading"
              @click="confirmRestore">
            <v-icon>
              mdi-backup-restore
            </v-icon>
          </v-btn>
        </v-layout>
      </v-col>
    </v-row>

    <div class="text-xs-center">
      <v-dialog v-model="dialog" width="400">
        <v-card>
          <v-card-title class="headline" primary-title>⚠️ 注意️</v-card-title>
          <v-card-text>{{message}}</v-card-text>
          <v-divider></v-divider>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" text @click="dialog = false; message = ''">好</v-btn>
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
        showPassword: false,
        loading: false,
        srcPath: "",
        password:"",
        rules: {
          required: value => !!value || '需要提供密码',
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
                self.message = error
                self.dialog = true
              }
            });
      },

      confirmRestore: function () {
        var self = this
        if (self.srcPath == "") {
          self.message = "请选择需要还原的文件！"
          self.dialog = true
        } else if (self.password == "" || self.password.length > 15) {
          self.message = "请正确输入密码！"
          self.dialog = true
        } else {
          self.loading = true
          window.backend.Backup.PerformRestore(self.srcPath, self.password).then(result => {
            self.message = result
            self.dialog = true
            self.loading = false
          })
          .catch(error => {
            self.message = error
            self.dialog = true
            self.loading = false
          });
        }
      },
    },
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
