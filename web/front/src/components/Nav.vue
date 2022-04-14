<template>
  <div>
    <v-card class="mx-auto" max-width="300">
      <v-img :src="profileInfo.img">
        <v-card-title>
          <v-col align="center">
            <v-avatar size="130" color="grey">
              <img :src="profileInfo.avatar" alt="" />
            </v-avatar>
            <div class="ma-4 white--text">{{profileInfo.name}}</div>
          </v-col>
        </v-card-title>
        <v-divider></v-divider>
        <!--分割线-->
      </v-img>

      <v-col align="center">
        <div class="ma-3">{{profileInfo.desc}}</div>
      </v-col>

      <v-divider></v-divider>
      <!--分割线-->

      <v-list nav dense>
        <v-list-item>
          <v-list-item-icon class="ma-3"
            ><v-icon>{{ "mdi-qqchat" }}</v-icon></v-list-item-icon
          >
          <v-list-item-content class="grey--text"
            >{{profileInfo.qq_chat}}</v-list-item-content
          >
        </v-list-item>

        <v-list-item>
          <v-list-item-icon class="ma-3"
            ><v-icon>{{ "mdi-email" }}</v-icon></v-list-item-icon
          >
          <v-list-item-content class="grey--text"
            >{{profileInfo.email}}</v-list-item-content
          >
        </v-list-item>

        <v-list-item>
          <v-list-item-icon class="ma-3"
            ><v-icon>{{ "mdi-github" }}</v-icon></v-list-item-icon
          >
          <v-list-item-content class="grey--text"
            >{{profileInfo.github}}</v-list-item-content
          >
        </v-list-item>


      </v-list>
    </v-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      profileInfo: {
        id: 1
      }
    }
  },
  created() {
    this.getProfileInfo()
  },
  methods: {
    // 获取个人设置
    async getProfileInfo() {
      const { data: res } = await this.$http.get(
        `profile/${this.profileInfo.id}`
      )
      this.profileInfo = res.data
    }
  }
}
</script>

<style></style>
