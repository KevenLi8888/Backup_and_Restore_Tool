<template>
  <v-container class="px-0">

    <v-row>
      <v-col>
        <v-card outlined class="pa-4 ma-4" height="85%">
          <v-card-title class="justify-center">
            选择需要备份的目录
          </v-card-title>
          <v-card-text class=text-center>
            {{srcPath}}
          </v-card-text>
          <v-card-actions>
            <v-layout justify-center align-center class="px-0">
              <v-btn color="#81D4FA" @click="selectSrcDir">打开...</v-btn>
            </v-layout>
          </v-card-actions>
        </v-card>
      </v-col>

      <v-col>
        <v-card outlined class="pa-4 ma-4" height="85%">
          <v-card-title class="justify-center">
            输入备份密码
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
        <v-card outlined class="pa-4 ma-4" height="85%">
          <v-card-title class="justify-center">
            选择目标备份路径
          </v-card-title>
          <v-card-text class=text-center v-if="this.destPath==''">
            不指定则默认备份至项目目录下./backup文件夹
          </v-card-text>
          <v-card-text class=text-center v-else>
            {{destPath}}
          </v-card-text>
          <v-card-actions>
            <v-layout justify-center align-center class="px-0">
              <v-btn color="#81D4FA" @click="selectDestDir">打开...</v-btn>
            </v-layout>
          </v-card-actions>
        </v-card>
      </v-col>

      <v-col>
        <v-card outlined class="pa-4 ma-4" height="85%">
          <v-card-title class="justify-center">
            输入备份文件名称
          </v-card-title>
          <v-layout justify-center align-center class="px-0">
            <v-text-field
                v-model="filename"
                label="输入文件名"
                outlined
                clearable
                hint="留空以使用默认名称"
                persistent-hint
                :rules="[rules.validFileName]"
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
              @click="confirmBackup">
            <v-icon>
              mdi-cloud-upload
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
        dialog: false,
        showPassword: false,
        loading: false,
        srcPath: "",
        destPath: "",
        password: "",
        filename: "",
        rules: {
          required: value => !!value || '需要设定密码',
          max: v => v.length < 16 || '至多输入15位密码',
          validFileName: value => {
            const pattern = /[/\\?%*:|"<>.？。：｜]/gi
            return !(pattern.test(value)) || '文件名含有非法字符'
          }
        },
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
                self.message = error
                self.dialog = true
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
                self.message = error
                self.dialog = true
              }
            });
      },

      confirmBackup: function () {
        var self = this
        const pattern = /[/\\?%*:|"<>.？。：｜]/gi
        if (self.srcPath == "") {
          self.message = "请选择需要备份的目录！"
          self.dialog = true
        } else if (self.password == "" || self.password.length > 15) {
          self.message = "请正确输入密码！"
          self.dialog = true
        } else if (pattern.test(self.filename)) {
          self.message = "请输入合法文件名！"
          self.dialog = true
        } else if (self.filename == "cache" || self.filename == "cache1") {
          self.message = "cache和cache1为系统保留名称，请更换文件名！"
          self.dialog = true
        } else if (self.srcPath == self.destPath) {
          self.message = "目标路径不能和备份目录相同，请重新选择目标路径！"
          self.dialog = true
        } else {
          self.loading = true
          window.backend.Backup.PerformBackup(self.srcPath, self.destPath, self.password, self.filename).then(result => {
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
