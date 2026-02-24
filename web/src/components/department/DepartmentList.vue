<script setup lang="ts">
import { h } from 'vue'
import { NButton, NPopconfirm } from 'naive-ui'
// hooks
import { useDeleteDepartmentMutation } from '@/data/department'
// utils
import { renderIcon } from '@/utils'
// types
import type { Category, Department } from '@/types'
// hooks
import { useModalStore } from '@/store/modal'
// components
import DepartmentModal from './DepartmentModal.vue'

// Define props
const props = defineProps<{
  loading: boolean
  tableData: Department[]
}>()
const modal = useModalStore()

// mutation
const { mutateAsync: deleteDepartment } = useDeleteDepartmentMutation()

function onEdit(department: Department) {
  modal.open(DepartmentModal, {
    title: 'Edit Department',
    props: {
      department,
    },
  })
}

async function deleteRow(row: Category) {
  await deleteDepartment({ id: row.id })
}

const columns = [
  {
    title: 'Name',
    key: 'name',
    width: 200,
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
            onClick: () => onEdit(row),
          },
          { default: () => 'Edit', icon: renderIcon('material-symbols:edit', { size: 16 }) },
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
                {
                  default: () => 'Delete',
                  icon: renderIcon('material-symbols:delete-outline', { size: 16 }),
                },
              ),
            default: () => 'Are you sure?',
          },
        ),
      ]
    },
  },
]
</script>

<template>
  <n-data-table :loading="loading" :columns="columns" :data="tableData" />
</template>
