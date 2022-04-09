<template>
  <div>
    <h3>用户列表</h3>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search
            v-model="queryParam.username"
            placeholder="输入查找用户名"
            enter-button
            allowClear
            @search="getUserList"
          />
        </a-col>
        <a-col :span="4">
          <a-button type="primary">新增用户</a-button>
        </a-col>
      </a-row>

      <a-table
        bordered
        rowKey="username"
        :columns="columns"
        :pagination="pagination"
        :dataSource="userlist"
        @change="handleTableChange"
      >
        <span slot="role" slot-scope="role">{{
          role == 1 ? '管理员' : '订阅者'
        }}</span>

        <template slot="action" slot-scope="data">
          <div class="ActionSlot">
            <a-button type="primary" style="margin-right:10px">编辑</a-button>
            <a-button type="danger" @click="deleteUser(data.ID)">删除</a-button>
          </div>
        </template>
      </a-table>
    </a-card>
  </div>
</template>

<script>
const columns = [
  //key只是标识符 dataindex是res中的命名
  {
    title: 'Id',
    key: 'id',
    dataIndex: 'ID',
    width: '10%',
    align: 'center'
  },
  {
    title: '用户名',
    key: 'username',
    dataIndex: 'username',
    width: '20%',
    align: 'center'
  },
  {
    title: '权限',
    dataIndex: 'role',
    key: 'role',
    width: '20%',
    align: 'center',
    scopedSlots: { customRender: 'role' }
  },
  {
    title: '操作',
    key: 'action',
    width: '20%',
    align: 'center',
    scopedSlots: { customRender: 'action' }
  }
]
export default {
  data() {
    return {
      pagination: {
        //分页对象
        pageSizeOptions: ['5', '10', '20'], //分页选项
        pageSize: 5, //默认一页几条
        total: 0, //数据总数
        showSizeChanger: true, //是否可以改变pagesize
        showTotal: total => `共 ${total} 条` //显示总共多少条
      },
      userlist: [],
      columns,
      queryParam: {
        username: '',
        pagesize: 5,
        pagenum: 1
      },
      visible: false
    }
  },
  created() {
    this.getUserList()
  },
  methods: {
    async getUserList() {
      const { data: res } = await this.$http.get('users', {
        //把params绑定到pagination对象
        params: {
          username: this.queryParam.username,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      console.log(res)
      if (res.status != 200) return this.$message.error(res.message) //返回报错
      this.pagination.total = res.total //设置总数
      this.userlist = res.data //设置
    },
    //分页
    handleTableChange(pagination, filters, sorter) {
      var pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.queryParam.pagesize = pagination.pageSize
      this.queryParam.pagenum = pagination.current

      if (pagination.pageSize !== this.pagination.pageSize) {
        this.queryParam.pagenum = 1
        pager.current = 1
      }
      this.pagination = pager
      this.getUserList()
    },
    //删除用户
     deleteUser(id) {
      this.$confirm({
        title: '提示：请再次确认',
        content: '确定要删除该用户吗？一旦删除，无法恢复',
        onOk: async () => {
          const res = await this.$http.delete(`user/${id}`)
          if (res.status != 200) return this.$message.error(res.message)
          this.$message.success('删除成功')
          this.getUserList()
        },
        onCancel: () => {
          this.$message.info('已取消删除')
        },
      })
    },
  },
}
</script>

<style scoped>
.ActionSlot {
  display: flex;
  justify-content: center;
}
</style>
