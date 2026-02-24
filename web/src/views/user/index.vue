<script setup lang="ts">
import { NButton } from 'naive-ui'
// hooks
import { useUsersQuery } from '@/data/user'
import { useModalStore } from '@/store/modal'
// components
import TheIcon from '@/components/icon/TheIcon.vue'
import UserList from '@/components/user/UserList.vue'
import CommonPage from '@/components/page/CommonPage.vue'
import UserCreateUpdateModal from '@/components/user/UserCreateUpdateModal.vue'

// query
const { users, loading } = useUsersQuery({})
// store hooks
const modal = useModalStore()

function openCreateModal() {
  modal.open(UserCreateUpdateModal, {
    title: 'Create User',
  })
}
</script>

<template>
  <CommonPage show-footer title="User List">
    <template #action>
      <div>
        <NButton class="float-right mr-15" type="primary" @click="openCreateModal">
          <TheIcon icon="material-symbols:add" :size="18" class="mr-5" />Create new
        </NButton>
      </div>
    </template>
    <UserList :loading="loading" :table-data="users" />
  </CommonPage>
</template>
