<template>
  <CommonPage show-footer title="Tasks">
    <template #action>
      <div>
        <NButton v-permission="'post/api/v1/dept/create'" class="float-right mr-15" type="primary">
          <TheIcon icon="material-symbols:add" :size="18" class="mr-5" />Create new task
        </NButton>
      </div>
    </template>
    <!-- <n-data-table
      :remote="remote"
      :loading="loading"
      :columns="columns"
      :data="tableData"
      :scroll-x="scrollX"
      :row-key="(row) => row[rowKey]"
      :pagination="isPagination ? pagination : false"
      @update:checked-row-keys="onChecked"
      @update:page="onPageChange"
    /> -->
  </CommonPage>
</template>

<script setup lang="ts">
import { h, onMounted, ref, resolveDirective, withDirectives } from 'vue'
import { NButton, NForm, NFormItem, NInput, NInputNumber, NPopconfirm, NTreeSelect } from 'naive-ui'
import CommonPage from '@/components/page/AppPage.vue'
import CrudModal from '@/components/table/CrudModal.vue'
import CrudTable from '@/components/table/CrudTable.vue'

const columns = [
  {
    title: 'Name',
    key: 'name',
    width: 200,
  },
  {
    title: 'Description',
    key: 'description',
    ellipsis: true,
    width: 300,
  },
  {
    title: 'Actions',
    key: 'actions',
    width: 160,
    render(row: any) {
      return [
        h(
          NButton,
          {
            size: 'small',
            type: 'primary',
            class: 'mr-5',
            onClick: () => editRow(row),
          },
          { default: () => 'Edit' },
        ),
        h(
          NPopconfirm,
          {
            onPositiveClick: () => deleteRow(row),
          },
          {
            trigger: () =>
              h(
                NButton,
                {
                  size: 'small',
                  type: 'error',
                },
                { default: () => 'Delete' },
              ),
            default: () => 'Are you sure?',
          },
        ),
      ]
    },
  },
]

function editRow(row: any) {
  console.log('Edit:', row)
}

function deleteRow(row: any) {
  console.log('Delete:', row)
}

const tableData = ref([
  { id: 1, name: 'Task A', description: 'Prepare documentation', loading: false },
  { id: 2, name: 'Task B', description: 'Build UI for dashboard', loading: false },
])
</script>
