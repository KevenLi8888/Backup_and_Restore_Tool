<template>
  <v-app id="inspire">

<!--    <v-toolbar-->
<!--        flat-->
<!--        color="primary"-->
<!--        dark-->
<!--    >-->
<!--      <v-toolbar-title>Backup & Restore Tool</v-toolbar-title>-->
<!--    </v-toolbar>-->
    <v-tabs
        v-model="tab"
        background-color="blue accent-2"
        centered
        dark
        icons-and-text
    >
      <v-tab>
        Backup
        <v-icon>
          upload_file
        </v-icon>
      </v-tab>
      <v-tab>
        Restore
        <v-icon>
          restore_page
        </v-icon>
      </v-tab>
      <v-tab>
        About
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
        <v-card flat>
          <v-card-text>
            <p>
              Fusce a quam. Phasellus nec sem in justo pellentesque facilisis. Nam eget dui. Proin viverra, ligula sit amet ultrices semper, ligula arcu tristique sapien, a accumsan nisi mauris ac eros. In dui magna, posuere eget, vestibulum et, tempor auctor, justo.
            </p>

            <p class="mb-0">
              Cras sagittis. Phasellus nec sem in justo pellentesque facilisis. Proin sapien ipsum, porta a, auctor quis, euismod ut, mi. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nam at tortor in tellus interdum sagittis.
            </p>
          </v-card-text>
        </v-card>
      </v-tab-item>
    </v-tabs>
  </v-app>
</template>

<script>
  import Backup from "./components/backup.vue"
  import Restore from "./components/restore.vue"
  import "material-design-icons-iconfont";



  export default {
    data: () => ({
      drawer: false
    }),
    components: {
      Backup,
      Restore
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
