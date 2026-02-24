<script setup lang="ts">
import { ref } from 'vue'
import { NButton } from 'naive-ui'
// hooks
import { useModalStore } from '@/store/modal'
import { useCategoriesQuery } from '@/data/category'
// components
import TheIcon from '@/components/icon/TheIcon.vue'
import CommonPage from '@/components/page/CommonPage.vue'
import CategoryList from '@/components/category/CategoryList.vue'
import CategoryModal from '@/components/category/CategoryModal.vue'

const page = ref(1)
const limit = ref(10)

// query
const { categories, paginatorInfo, loading } = useCategoriesQuery({
  limit: limit,
  page: page,
})

function handlePageChange(newPage: number) {
  page.value = newPage
}

function handlePageSizeChange(newSize: number) {
  limit.value = newSize
  page.value = 1 // reset to first page (important UX rule)
}

// store hooks
const modal = useModalStore()

function openCreateModal() {
  modal.open(CategoryModal, {
    title: 'Create Category',
  })
}
</script>

<template>
  <CommonPage show-footer title="Category List">
    <template #action>
      <div>
        <NButton class="float-right mr-15" type="primary" @click="openCreateModal">
          <TheIcon icon="material-symbols:add" :size="18" class="mr-5" />Create new
        </NButton>
      </div>
    </template>
    <CategoryList
      :loading="loading"
      :table-data="categories"
      :paginatorInfo="paginatorInfo"
      :page="page"
      :limit="limit"
      @update:page="handlePageChange"
      @update:pageSize="handlePageSizeChange"
    />
  </CommonPage>
</template>
