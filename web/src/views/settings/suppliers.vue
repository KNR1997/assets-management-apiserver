<script setup lang="ts">
import { NButton } from 'naive-ui'
// hooks
import { useModalStore } from '@/store/modal'
import { useSuppliersQuery } from '@/data/supplier'
// components
import TheIcon from '@/components/icon/TheIcon.vue'
import CommonPage from '@/components/page/CommonPage.vue'
import SupplierModal from '@/components/supplier/SupplierModal.vue'
import SupplierList from '@/components/supplier/SupplierList.vue'

// query
const { suppliers, loading } = useSuppliersQuery({})
// store hooks
const modal = useModalStore()

function openCreateModal() {
  modal.open(SupplierModal, {
    title: 'Create Supplier',
  })
}
</script>

<template>
  <CommonPage show-footer title="Supplier List">
    <template #action>
      <div>
        <NButton class="float-right mr-15" type="primary" @click="openCreateModal">
          <TheIcon icon="material-symbols:add" :size="18" class="mr-5" />Create new
        </NButton>
      </div>
    </template>
    <SupplierList :loading="loading" :table-data="suppliers" />
  </CommonPage>
</template>
