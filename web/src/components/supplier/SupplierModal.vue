<script setup lang="ts">
import { ref, watch } from 'vue'
import { NForm, NFormItem, NInput, NButton } from 'naive-ui'
// hooks
import { useModalStore } from '@/store/modal'
import { useCreateSupplierMutation, useUpdateSupplierMutation } from '@/data/supplier';
// types
import type { Supplier } from '@/types'

const props = defineProps<{
  supplier?: Supplier | null
}>()

const modal = useModalStore()

// mutations
const { mutateAsync: createSupplier, isPending: creating } = useCreateSupplierMutation()
const { mutateAsync: updateSupplier, isPending: updating } = useUpdateSupplierMutation()

const modalFormRef = ref()
const modalForm = ref({
  name: '',
  email: '',
})

watch(
  () => props.supplier,
  (supplier) => {
    if (!supplier) {
      // create mode
      modalForm.value = {
        name: '',
        email: '',
      }
      return
    }

    // edit mode
    modalForm.value = {
      name: supplier.name,
      email: supplier.email,
    }
  },
  { immediate: true },
)

const validationRules = {
  name: [{ required: true, message: 'Name is required', trigger: ['blur'] }],
  // email: [{ required: true, message: 'Email is required', trigger: ['blur'] }],
}

async function handleSave() {
  modalFormRef.value?.validate(async (errors: Error) => {
    if (errors) return
    if (props.supplier) {
      await updateSupplier({
        id: props.supplier.id,
        name: modalForm.value.name,
        // email: modalForm.value.email,
      })
    } else {
      await createSupplier({
        name: modalForm.value.name,
        // email: modalForm.value.email,
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
      <NFormItem label="Name" path="name">
        <NInput v-model:value="modalForm.name" />
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
