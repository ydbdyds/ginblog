<template>
  <v-app app>
    <TopBar></TopBar>
    <v-sheet
      id="scrolling-techniques-3"
      class="overflow-y-auto"
      max-height="950"
    >
      <v-main class="grey lighten-3">
        <v-container>
          <v-row>
            <v-col cols="3"><Live2d></Live2d> </v-col>
            <v-col>
              <v-sheet max-width="65vw" min-height="80vh" rounded="lg">
                <router-view :key="$route.path"></router-view>
              </v-sheet>
            </v-col>
            <v-col cols="3">
              <Nav></Nav>
            </v-col>
          </v-row>
        </v-container>
      </v-main>
    </v-sheet>
    <Footer></Footer>
  </v-app>
</template>

<script>
import TopBar from "../components/TopBar";
import Footer from "../components/Footer";
import Nav from "../components/Nav";
import Live2d from "../components/Live2d/index";
export default {
  components: { TopBar, Footer, Nav, Live2d },
};
</script>
