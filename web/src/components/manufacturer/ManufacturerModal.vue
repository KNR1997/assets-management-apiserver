<script setup lang="ts">
import { ref, watch } from 'vue'
import { NForm, NFormItem, NInput, NButton } from 'naive-ui'
// hooks
import { useModalStore } from '@/store/modal'
import { useCreateManufacturerMutation, useUpdateManufacturerMutation } from '@/data/manufacturer'
// types
import type { Manufacturer } from '@/types'

const props = defineProps<{
  manufacturer?: Manufacturer | null
}>()

const modal = useModalStore()

// mutations
const { mutateAsync: createManufacturer, isPending: creating } = useCreateManufacturerMutation()
const { mutateAsync: updateManufacturer, isPending: updating } = useUpdateManufacturerMutation()

const modalFormRef = ref()
const modalForm = ref({
  name: '',
  email: '',
})

watch(
  () => props.manufacturer,
  (manufacturer) => {
    if (!manufacturer) {
      // create mode
      modalForm.value = {
        name: '',
        email: '',
      }
      return
    }

    // edit mode
    modalForm.value = {
      name: manufacturer.name,
      email: manufacturer.email,
    }
  },
  { immediate: true },
)

const validationRules = {
  name: [{ required: true, message: 'Name is required', trigger: ['blur'] }],
  email: [{ required: true, message: 'Email is required', trigger: ['blur'] }],
}

async function handleSave() {
  modalFormRef.value?.validate(async (errors: Error) => {
    if (errors) return
    if (props.manufacturer) {
      await updateManufacturer({
        id: props.manufacturer.id,
        name: modalForm.value.name,
        email: modalForm.value.email,
      })
    } else {
      await createManufacturer({
        name: modalForm.value.name,
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
