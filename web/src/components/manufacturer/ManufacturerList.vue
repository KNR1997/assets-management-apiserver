<script setup lang="ts">
import { h } from 'vue'
import { NButton, NPopconfirm } from 'naive-ui'
// hooks
import { useDeleteManufacturerMutation } from '@/data/manufacturer'
// utils
import { renderIcon } from '@/utils'
// types
import type { Category, Manufacturer } from '@/types'
// hooks
import { useModalStore } from '@/store/modal'
// components
import ManufacturerModal from './ManufacturerModal.vue'

// Define props
const props = defineProps<{
  loading: boolean
  tableData: Manufacturer[]
}>()
const modal = useModalStore()

// mutation
const { mutateAsync: deleteManufacturer } = useDeleteManufacturerMutation()

function onEdit(manufacturer: Manufacturer) {
  modal.open(ManufacturerModal, {
    title: 'Edit Manufacturer',
    props: {
      manufacturer,
    },
  })
}

async function deleteRow(row: Category) {
  await deleteManufacturer({ id: row.id })
}

const columns = [
  {
    title: 'Name',
    key: 'name',
    width: 200,
  },
  {
    title: 'Email',
    key: 'email',
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
