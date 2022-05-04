<template>
  <v-col>
    <div v-if="total == 0 && isLoad" class="d-flex justify-center align-center">
      <div>
        <v-alert class="ma-5" dense outlined type="error">抱歉，暂无数据！</v-alert>
      </div>
    </div>
    <v-sheet>
      <v-card
        class="ma-3"
        v-for="item in artList"
        :key="item.id"
        link
        @click="$router.push(`/detail/${item.ID}`)"
      >
        <v-row no-gutters>
        <v-avatar class="ma-3 hidden-sm-and-down" size="125" tile>
          <v-img :src="item.img"></v-img>
        </v-avatar>
        <v-col>
          <v-card-title class="my-2"
            ><v-chip color="pink" label class="mr-3 white--text"
              ><v-icon left> mdi-label </v-icon>{{ item.Category.name }}</v-chip
            >{{ item.title }}</v-card-title
          >
          <v-card-subtitle>{{ item.desc }}</v-card-subtitle>
          <v-divider></v-divider>
          <v-card-text>
            <v-icon class="mr-1" small>{{ "mdi-calendar-month" }}</v-icon>
            <span class="mr-10">{{ item.CreatedAt | dateformat("YYYY-MM-DD HH:SS") }}</span>
            <v-icon class="mr-1" small>{{ "mdi-fire" }}</v-icon>
            <span>{{ item.view }}</span>
          </v-card-text>
        </v-col>
      </v-row>
      </v-card>
      <v-col>
        <div class="text-center">
          <v-pagination
            total-visible="7"
            v-model="queryParam.pagenum"
            :length="Math.ceil(total/queryParam.pagesize)"
            @input="getArtList()"
          ></v-pagination>
        </div>
      </v-col>
    </v-sheet>
  </v-col>
</template>
<script>
export default {
  props: ['cid'],
  data() {
    return {
      artList: [],
      queryParam: {
        pagesize: 5,
        pagenum: 1
      },
      total: 0,
      isLoad: false
    }
  },
  mounted() {
    this.getArtList()
  },
  methods: {
    // 获取文章列表
    async getArtList() {
      const { data: res } = await this.$http.get(`article/list/${this.cid}`, {
        params: {
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })

      this.artList = res.data
      this.total = res.total
      this.isLoad = true
    }
  }
}
</script>
<style scoped>
.nodate {
  width: 100%;
  height: 100%;
}
</style>