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
          <a-button type="primary" @click="addUserVisible = true"
            >新增用户</a-button
          >
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
            <a-button
              icon = "edit"
              type="primary"
              style="margin-right:15px"
              @click="editUser(data.ID)"
              >编辑</a-button
            >
            <a-button type="danger" icon="delete" style="margin-right:15px" @click="deleteUser(data.ID)">删除</a-button>
             <a-button type="info" icon="rest" @click="deleteUser(data.ID)">重置密码</a-button>
          </div>
        </template>
      </a-table>
    </a-card>

    <!-- 新增用户 -->
    <a-modal
      closable
      title="新增用户"
      :visible="addUserVisible"
      width="60%"
      @ok="addUserOk"
      @cancel="addUserCancel"
      :destroyOnClose="true"
    >
      <a-form-model :model="newUser" :rules="addUserRules" ref="addUserRef">
        <a-form-model-item label="用户名" prop="username">
          <a-input v-model="newUser.username"></a-input>
        </a-form-model-item>
        <a-form-model-item has-feedback label="密码" prop="password">
          <a-input-password v-model="newUser.password"></a-input-password>
        </a-form-model-item>
        <a-form-model-item has-feedback label="确认密码" prop="checkpass">
          <a-input-password v-model="newUser.checkpass"></a-input-password>
        </a-form-model-item>
      </a-form-model>
    </a-modal>

    <!-- 编辑用户 -->
    <a-modal
      closable
      title="编辑用户"
      :visible="editUserVisible"
      width="60%"
      @ok="editUserOk"
      @cancel="editUserCancel"
    >
      <a-form-model :model="userInfo" :rules="userRules" ref="addUserRef">
        <a-form-model-item label="用户名" prop="username">
          <a-input v-model="userInfo.username"></a-input>
        </a-form-model-item>
        <a-form-model-item label="是否为管理员">
          <a-switch
            checked-children="是"
            un-checked-children="否"
            @change="adminChange"
          />
        </a-form-model-item>
      </a-form-model>
    </a-modal>
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
      userInfo: {
        id: 0,
        username: '',
        password: '',
        checkpass: '',
        role: 2
      },
      newUser: {
        id: 0,
        username: '',
        password: '',
        checkpass: '',
        role: 2
      },
      visible: false,
      addUserVisible: false,
      editUserVisible: false,
      userRules: {
        //输入规则判断
        username: [
          {
            validator: (rule, value, callback) => {
              if (this.userInfo.username == '') {
                callback(new Error('请输入用户名'))
              }
              if (
                [...this.userInfo.username].length < 4 ||
                [...this.userInfo.username].length > 12
              ) {
                callback(new Error('用户名应当在4到12个字符之间'))
              } else {
                callback()
              }
            },
            trigger: 'blur' //触发机制
          }
        ],
        password: [
          {
            validator: (rule, value, callback) => {
              if (this.userInfo.password == '') {
                callback(new Error('请输入密码'))
              }
              if (
                [...this.userInfo.password].length < 6 ||
                [...this.userInfo.password].length > 20
              ) {
                callback(new Error('密码应当在6到20位之间'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        checkpass: [
          {
            validator: (rule, value, callback) => {
              if (this.userInfo.checkpass == '') {
                callback(new Error('请输入密码'))
              }
              if (this.userInfo.password !== this.userInfo.checkpass) {
                callback(new Error('密码不一致，请重新输入'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      },
      addUserRules: {
        //输入规则判断
        username: [
          {
            validator: (rule, value, callback) => {
              if (this.newUser.username == '') {
                callback(new Error('请输入用户名'))
              }
              if (
                [...this.newUser.username].length < 4 ||
                [...this.newUser.username].length > 12
              ) {
                callback(new Error('用户名应当在4到12个字符之间'))
              } else {
                callback()
              }
            },
            trigger: 'blur' //触发机制
          }
        ],
        password: [
          {
            validator: (rule, value, callback) => {
              if (this.newUser.password == '') {
                callback(new Error('请输入密码'))
              }
              if (
                [...this.newUser.password].length < 6 ||
                [...this.newUser.password].length > 20
              ) {
                callback(new Error('密码应当在6到20位之间'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        checkpass: [
          {
            validator: (rule, value, callback) => {
              if (this.newUser.checkpass == '') {
                callback(new Error('请输入密码'))
              }
              if (this.newUser.password !== this.newUser.checkpass) {
                callback(new Error('密码不一致，请重新输入'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      }
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
      //console.log(res)
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
        }
      })
    },
    // 新增用户
    addUserOk() {
      this.$refs.addUserRef.validate(async valid => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.post('user/add', {
          username: this.newUser.username,
          password: this.newUser.password,
          role: this.newUser.role
        })
        if (res.status != 200) return this.$message.error(res.message)
        this.addUserVisible = false
        this.$message.success('添加用户成功')
        this.getUserList()
      })
    },
    addUserCancel() {
      this.$refs.addUserRef.resetFields() //清空
      this.addUserVisible = false //重置
      this.$message.info('编辑已取消')
    },
    adminChange(value) {
      //选择框发生变动执行
      this.userInfo.role = Number(value)
      //console.log(this.userInfo.role)
    },
    // 编辑用户
    async editUser(id) {
      this.editUserVisible = true
      const { data: res } = await this.$http.get(`user/${id}`)
      this.userInfo = res.data
      this.userInfo.id = id
    },
    editUserOk() {
      this.$refs.addUserRef.validate(async valid => {
        if (!valid) return this.$message.error('参数不符合要求，请重新输入')
        const { data: res } = await this.$http.put(`user/${this.userInfo.id}`, {
          username: this.userInfo.username,
          role: this.userInfo.role
        })
        if (res.status != 200) return this.$message.error(res.message)
        this.editUserVisible = false
        this.$message.success('更新用户信息成功')
        this.getUserList()
      })
    },
    editUserCancel() {
      this.$refs.addUserRef.resetFields()
      this.editUserVisible = false
      this.$message.info('编辑已取消')
    }
  }
}
</script>

<style scoped>
.ActionSlot {
  display: flex;
  justify-content: center;
}
</style>
