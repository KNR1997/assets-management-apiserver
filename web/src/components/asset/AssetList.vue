<script setup lang="ts">
import { h } from 'vue'
import { NButton, NPopconfirm } from 'naive-ui'
// hooks
import { useDeleteAssetMutation } from '@/data/asset'
// utils
import { renderIcon } from '@/utils'
// types
import type { Asset } from '@/types'
// hooks
import { useModalStore } from '@/store/modal'
// components
import AssetModal from './AssetModal.vue'

// Define props
const props = defineProps<{
  loading: boolean
  tableData: Asset[]
}>()
const modal = useModalStore()

// mutation
const { mutateAsync: deleteAsset } = useDeleteAssetMutation()

function onEdit(asset: any) {
  modal.open(AssetModal, {
    title: 'Edit Asset',
    props: {
      asset,
    },
  })
}

async function deleteRow(row: any) {
  await deleteAsset({ id: row.ID })
}

const columns = [
  {
    title: 'Name',
    key: 'name',
    width: 200,
  },
  {
    title: 'Serial No.',
    key: 'serialNumber',
    ellipsis: true,
    width: 300,
  },
    {
    title: 'Status',
    key: 'status',
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
