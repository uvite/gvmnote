<template>
  <div>
    <warning-bar href="https://www.bilibili.com/video/BV1kv4y1g7nT?p=3"
                 title="此功能为开发环境使用，不建议发布到生产，具体使用效果请看视频https://www.bilibili.com/video/BV1kv4y1g7nT?p=3"/>
    <div class="gva-table-box">
      <div class="gva-btn-list">

        <a-space>

          <a-button type="primary" @click="openDialog('addApi')">
            <template #icon>
              <icon-plus/>
            </template>
            <!-- Use the default slot to avoid extra spaces -->
            <template #default>新增</template>
          </a-button>
        </a-space>
      </div>


      <a-table :data="tableData" style="margin-top: 30px">
        <template #columns>

          <a-table-column title="id" data-index="id"></a-table-column>
          <a-table-column title="包名" data-index="packageName"></a-table-column>
          <a-table-column title="展示名" data-index="label"></a-table-column>
          <a-table-column title="描述" data-index="desc"></a-table-column>
          <a-table-column title="Optional">


            <template #cell="{ scope }">

              <a-popconfirm
                  content="确定要删除该数据吗?"
                  position="bottom"
                  @ok="deleteApiFunc(scope.row)"

              >
                <a-link
                    type="primary"
                    v-auth="options.delete.auth || []"
                    v-role="options.delete.role || []"
                >
                  <icon-delete/>
                  删除
                </a-link>
              </a-popconfirm>


            </template>
          </a-table-column>
        </template>
      </a-table>


    </div>

    <a-modal v-model:visible="dialogFormVisible" title="Modal Form" @cancel="closeDialog" @before-ok="enterDialog">
      <a-form ref="pkgForm" :model="form" :rules="rules" label-width="80px">
        <a-form-item field="packageName" label="包名"
                     :rules="[{required:true,message:'packageName is required'},{minLength:1,message:'must be greater than 1 characters'}]"
                     :validate-trigger="['change','input']">
          <a-input v-model="form.packageName"/>
        </a-form-item>
        <a-form-item field="label" label="展示名">
          <a-input v-model="form.label"/>
        </a-form-item>
        <a-form-item field="desc" label="描述">
          <a-input
              v-model="form.desc"
              placeholder=""
          />
        </a-form-item>
      </a-form>
    </a-modal>


  </div>
</template>

<script>
export default {
  name: 'AutoPkg',
}
</script>

<script setup>
import {
  createPackageApi,
  getPackageApi,
  deletePackageApi,
} from '@/api/tools/autoCode'


import {ref} from 'vue'
// import WarningBar from '@/components/warningBar/warningBar.vue'
// import { ElMessage, ElMessageBox } from 'element-plus'
import {Message} from '@arco-design/web-vue'

const form = ref({
  packageName: '',
  label: '',
  desc: '',
})


const dialogFormVisible = ref(false)
const openDialog = () => {
  dialogFormVisible.value = true
}

const closeDialog = () => {
  dialogFormVisible.value = false
  form.value = {
    packageName: '',
    label: '',
    desc: '',
  }
}

const pkgForm = ref(null)
const enterDialog = async () => {

  const res = await createPackageApi(form.value)
  if (res.code === 0) {
    Message.success('添加成功')

  }
  getTableData()
  closeDialog()
    

}

const tableData = ref([])
const getTableData = async () => {
  const table = await getPackageApi()
  if (table.code === 0) {
    tableData.value = table.data.pkgs
  }
}

const deleteApiFunc = async (row) => {
  async () => {
    const res = await deletePackageApi(row)
    if (res.code === 0) {
      Message.success('删除成功')

      getTableData()
    }
  }
}

getTableData()
</script>


