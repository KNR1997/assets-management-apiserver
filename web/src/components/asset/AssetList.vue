<script setup lang="ts">
import { h } from 'vue'
import { NButton, NPopconfirm, NTag } from 'naive-ui'
// hooks
import { useDeleteAssetMutation } from '@/data/asset'
// utils
import { renderIcon } from '@/utils'
// types
import type { Asset } from '@/types'
import { AssetStatus } from '@/types/enum'
// hooks
import { useRouter } from 'vue-router'
import { useModalStore } from '@/store/modal'
// components
import CheckoutModal from './CheckoutModal.vue'
import CheckinModal from './CheckinModal.vue'

// Define props
const props = defineProps<{
  loading: boolean
  tableData: Asset[]
}>()
const modal = useModalStore()
const router = useRouter()
// mutation
const { mutateAsync: deleteAsset } = useDeleteAssetMutation()

function onEdit(asset: Asset) {
  router.push({
    name: 'Asset Edit',
    params: {
      id: asset.id,
    },
  })
}

function handleCheckoutCheckin(asset: Asset) {
  if (asset.status == AssetStatus.AssetAssigned) {
    modal.open(CheckinModal, {
      title: 'Checkin Asset',
      props: {
        asset,
      },
    })
  } else if (asset.status == AssetStatus.AssetAvailable) {
    modal.open(CheckoutModal, {
      title: 'Checkout Asset',
      props: {
        asset,
      },
    })
  }
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
    title: 'Serial',
    key: 'serialNumber',
    ellipsis: true,
    width: 150,
  },
  // {
  //   title: 'Model',
  //   key: 'model',
  //   ellipsis: true,
  //   width: 150,
  // },
  {
    title: 'Status',
    key: 'status',
    width: 150,
    render(row: Asset) {
      const isDeployed = row.status === AssetStatus.AssetAssigned
      return h(
        NTag,
        {
          round: true,
          bordered: false,
          type: isDeployed ? 'warning' : 'success',
        },
        {
          default: () => (isDeployed ? 'Deployed' : row.status),
        },
      )
    },
  },
  {
    title: 'Checkin/Checkout',
    key: 'status',
    width: 200,
    render(row: Asset) {
      return [
        h(
          NButton,
          {
            size: 'small',
            color: row.status == AssetStatus.AssetAssigned ? '#605ca8' : '#D81B60',
            type: 'primary',
            onClick: () => handleCheckoutCheckin(row),
            disabled: row.status == AssetStatus.AssetBroken,
          },
          { default: () => (row.status == AssetStatus.AssetAssigned ? 'Checkin' : 'Checkout') },
        ),
      ]
    },
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
