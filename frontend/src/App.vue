<template>
  <v-app id="inspire">
    <v-navigation-drawer v-model="drawer" clipped fixed app>
      <v-list dense>
        <v-list-item>
          <v-list-item-icon>
            <v-icon>upload_file</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>Backup Files</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-icon>
            <v-icon>restore_page</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>Restore Files</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-app-bar app fixed clipped-left>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title>Backup Tool</v-toolbar-title>
    </v-app-bar>
    <v-main>
      <v-container fluid class="px-0">
        <v-layout justify-center align-center class="px-0">
          <back-up></back-up>
        </v-layout>
      </v-container>
    </v-main>
    <v-footer app fixed>
      <span style="margin-left:1em">&copy; You</span>
    </v-footer>
  </v-app>
</template>

<script>
  import BackUp from "./components/backup.vue"
  import "material-design-icons-iconfont";


  export default {
    data: () => ({
      drawer: true
    }),
    components: {
      BackUp
    },
    props: {
      source: String
    }
  }
</script>

<style>
  .logo {
    width: 16em;
  }
</style>
