<template>
  <v-app id="inspire">
    <v-tabs
        v-model="tab"
        background-color="blue accent-2"
        centered
        dark
        icons-and-text
    >
      <v-tab>
        备份
        <v-icon>
          upload_file
        </v-icon>
      </v-tab>
      <v-tab>
        还原
        <v-icon>
          restore_page
        </v-icon>
      </v-tab>
      <v-tab>
        关于
        <v-icon>
          info
        </v-icon>
      </v-tab>

      <v-tab-item>
        <v-container fluid class="px-0">
          <v-layout justify-center align-center class="px-0">
            <backup></backup>
          </v-layout>
        </v-container>
      </v-tab-item>
      <v-tab-item>
        <v-container fluid class="px-0">
          <v-layout justify-center align-center class="px-0">
            <restore></restore>
          </v-layout>
        </v-container>
      </v-tab-item>
      <v-tab-item>
        <v-container fluid class="px-0">
          <v-layout justify-center align-center class="px-0">
            <about></about>
          </v-layout>
        </v-container>
      </v-tab-item>
    </v-tabs>
  </v-app>
</template>

<script>
  import Backup from "./components/backup.vue"
  import Restore from "./components/restore.vue"
  import About from "./components/about.vue"
  import "material-design-icons-iconfont";



  export default {
    data: () => ({
      drawer: false
    }),
    components: {
      Backup,
      Restore,
      About
    },
    props: {
      source: String
    }
  }
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Noto+Sans+SC:wght@300;400;500;700&display=swap');


.v-application {
  font-family: 'Noto Sans SC', sans-serif}

.logo {
    width: 16em;
  }

  h1 {
    margin-top: 2em;
    position: relative;
    min-height: 5rem;
    width: 100%;
  }

</style>
