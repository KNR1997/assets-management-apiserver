<script setup lang="ts">
import { NButton } from 'naive-ui'
// hooks
import { useModalStore } from '@/store/modal'
import { useManufacturersQuery } from '@/data/manufacturer'
// components
import TheIcon from '@/components/icon/TheIcon.vue'
import CommonPage from '@/components/page/CommonPage.vue'
import ManufacturerList from '@/components/manufacturer/ManufacturerList.vue'
import ManufacturerModal from '@/components/manufacturer/ManufacturerModal.vue'

// query
const { manufacturers, loading } = useManufacturersQuery({})
// store hooks
const modal = useModalStore()

function openCreateModal() {
  modal.open(ManufacturerModal, {
    title: 'Create Manufacturer',
  })
}
</script>

<template>
  <CommonPage show-footer title="Manufacturer List">
    <template #action>
      <div>
        <NButton class="float-right mr-15" type="primary" @click="openCreateModal">
          <TheIcon icon="material-symbols:add" :size="18" class="mr-5" />Create new
        </NButton>
      </div>
    </template>
    <ManufacturerList :loading="loading" :table-data="manufacturers" />
  </CommonPage>
</template>
