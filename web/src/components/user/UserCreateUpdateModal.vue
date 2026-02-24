<script setup lang="ts">
import { ref, watch } from 'vue'
import { NForm, NFormItem, NInput, NButton } from 'naive-ui'
// hooks
import { useModalStore } from '@/store/modal'
import { useCreateCategoryMutation, useUpdateCategoryMutation } from '@/data/category'
// types
import type { Category, User } from '@/types'
import { useCreateUserMutation, useUpdateUserMutation } from '@/data/user'

const props = defineProps<{
  user?: User | null
}>()

const modal = useModalStore()

// mutations
const { mutateAsync: createUser, isPending: creating } = useCreateUserMutation()
const { mutateAsync: updateUser, isPending: updating } = useUpdateUserMutation()

const modalFormRef = ref()
const modalForm = ref({
  username: '',
  email: '',
})

watch(
  () => props.user,
  (user) => {
    if (!user) {
      // create mode
      modalForm.value = {
        username: '',
        email: ''
      }
      return
    }

    // edit mode
    modalForm.value = {
      username: user.username,
      email: user.email,
    }
  },
  { immediate: true },
)

const validationRules = {
  username: [{ required: true, message: 'Username is required', trigger: ['blur'] }],
  email: [{ required: true, message: 'Email is required', trigger: ['blur'] }],
}

async function handleSave() {
  modalFormRef.value?.validate(async (errors: Error) => {
    if (errors) return
    if (props.user) {
      await updateUser({
        id: props.user.id,
        username: modalForm.value.username,
        email: modalForm.value.email,
      })
    } else {
      await createUser({
        username: modalForm.value.username,
        email: modalForm.value.email,
      })
    }
    modal.close()
  })
}
</script>

<template>
  <div>
    <!-- FORM -->
    <NForm
      ref="modalFormRef"
      label-placement="left"
      label-align="left"
      :label-width="80"
      :model="modalForm"
      :rules="validationRules"
    >
      <NFormItem label="Username" path="username">
        <NInput v-model:value="modalForm.username" />
      </NFormItem>

      <NFormItem label="Email" path="email">
        <NInput v-model:value="modalForm.email" />
      </NFormItem>
    </NForm>

    <div flex justify-end>
      <NButton @click="modal.close">Cancel</NButton>
      <NButton type="primary" class="ml-16" :loading="creating || updating" @click="handleSave">
        Save
      </NButton>
    </div>
  </div>
</template>
