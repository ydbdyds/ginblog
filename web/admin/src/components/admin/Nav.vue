<template>
  <a-layout-sider breakpoint="lg" v-model="collapsed">
    <div class="logo">
      <span>{{ collapsed ? '博客' : '风起的博客'}}</span>
    </div>

    <a-menu
      theme="dark"
      mode="inline"
      @click="goToPage"
      :default-selected-keys="['index']"
    >
      <a-menu-item key="index"
        ><a-icon type="dashboard"></a-icon><span>仪表盘</span></a-menu-item
      >
      <a-sub-menu>
        <span slot="title"><a-icon type="file" /><span>文章管理</span></span>
        <a-menu-item key="addart"
          ><a-icon type="form" /><span>写文章</span></a-menu-item
        >
        <a-menu-item key="artlist"
          ><a-icon type="snippets" /><span>文章列表</span></a-menu-item
        >
      </a-sub-menu>
      <a-menu-item key="catelist"
        ><a-icon type="book" /><span>分类列表</span></a-menu-item
      >
      <a-menu-item key="userlist"
        ><a-icon type="user" /><span>用户列表</span></a-menu-item
      >
    </a-menu>
  </a-layout-sider>
</template>

<script>
export default {
  data() {
    return {
      collapsed: false
    }
  },
    methods: {
    goToPage(item) {
      this.$router.push('/admin/' + item.key).catch((err) => err)
    },
  },
}
</script>


<style scoped>
.logo {
  height: 32px;
  margin: 16px;
  background-color: white;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 17px;
}
</style>
