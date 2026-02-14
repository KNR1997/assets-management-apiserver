<script setup lang="ts">
import { NButton } from 'naive-ui'
// hooks
import { useModalStore } from '@/store/modal'
import { useAssetsQuery } from '@/data/asset'
// components
import TheIcon from '@/components/icon/TheIcon.vue'
import CommonPage from '@/components/page/CommonPage.vue'
import AssetList from '@/components/asset/AssetList.vue'
import AssetModal from '@/components/asset/AssetModal.vue'

// query
const { assets, loading } = useAssetsQuery({})
// store hooks
const modal = useModalStore()

function openCreateModal() {
  modal.open(AssetModal, {
    title: 'Create Asset',
  })
}
</script>

<template>
  <CommonPage show-footer title="Asset List">
    <template #action>
      <div>
        <NButton class="float-right mr-15" type="primary" @click="openCreateModal">
          <TheIcon icon="material-symbols:add" :size="18" class="mr-5" />Create new asset
        </NButton>
      </div>
    </template>
    <AssetList :loading="loading" :table-data="assets" />
  </CommonPage>
</template>
