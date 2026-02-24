<script setup lang="ts">
import { ref, watch } from 'vue'
import { NForm, NFormItem, NInput, NButton } from 'naive-ui'
// hooks
import { useModalStore } from '@/store/modal'
import { useCreateDepartmentMutation, useUpdateDepartmentMutation } from '@/data/department'
// types
import type { Department } from '@/types'

const props = defineProps<{
  department?: Department | null
}>()

const modal = useModalStore()

// mutations
const { mutateAsync: createDepartment, isPending: creating } = useCreateDepartmentMutation()
const { mutateAsync: updateDepartment, isPending: updating } = useUpdateDepartmentMutation()

const modalFormRef = ref()
const modalForm = ref({
  name: '',
  notes: '',
})

watch(
  () => props.department,
  (department) => {
    if (!department) {
      // create mode
      modalForm.value = {
        name: '',
        notes: '',
      }
      return
    }

    // edit mode
    modalForm.value = {
      name: department.name,
      notes: department.notes,
    }
  },
  { immediate: true },
)

const validationRules = {
  name: [{ required: true, message: 'Name is required', trigger: ['blur'] }],
}

async function handleSave() {
  modalFormRef.value?.validate(async (errors: Error) => {
    if (errors) return
    if (props.department) {
      await updateDepartment({
        id: props.department.id,
        name: modalForm.value.name,
        notes: modalForm.value.notes,
      })
    } else {
      await createDepartment({
        name: modalForm.value.name,
        notes: modalForm.value.notes,
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

      <NFormItem label="Notes" path="notes">
        <NInput type="textarea" v-model:value="modalForm.notes" />
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
