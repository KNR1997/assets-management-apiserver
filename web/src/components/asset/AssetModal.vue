<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { NForm, NFormItem, NInput, NButton } from 'naive-ui'
// hooks
import { useModalStore } from '@/store/modal'
import { useCategoriesQuery } from '@/data/category'
import { useCreateAssetMutation, useUpdateAssetMutation } from '@/data/asset'
// types
import type { Asset } from '@/types'

const props = defineProps<{
  asset?: Asset | null
}>()

const modal = useModalStore()

// query
const { categories, loading: categoriesLoading } = useCategoriesQuery({})

// mutations
const { mutateAsync: createAsset, isPending: creating } = useCreateAssetMutation()
const { mutateAsync: updateAsset, isPending: updating } = useUpdateAssetMutation()

const modalFormRef = ref()
const modalForm = ref({
  name: '',
  serialNumber: '',
  categoryId: '',
})

watch(
  () => props.asset,
  (asset) => {
    if (!asset) {
      // create mode
      modalForm.value = {
        name: '',
        serialNumber: '',
        categoryId: '',
      }
      return
    }

    // edit mode
    modalForm.value = {
      name: asset.name,
      serialNumber: asset.serialNumber,
      categoryId: asset.categoryId,
    }
  },
  { immediate: true },
)

const categoryOptions = computed(
  () =>
    categories.value?.map((category) => ({
      label: category.name,
      value: category.id,
    })) ?? [],
)

const validateAddCourse = {
  name: [{ required: true, message: 'Name is required', trigger: ['blur'] }],
  serialNumber: [{ required: true, message: 'Serial Number is required', trigger: ['blur'] }],
}

async function handleSave() {
  modalFormRef.value?.validate(async (errors: Error) => {
    if (errors) return
    if (props.asset) {
      await updateAsset({
        id: props.asset.id,
        name: modalForm.value.name,
        // serialNumber: modalForm.value.serialNumber,
      })
    } else {
      await createAsset({
        name: modalForm.value.name,
        serialNumber: modalForm.value.serialNumber,
        description: null,
        categoryID: 1,
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
      :rules="validateAddCourse"
    >
      <NFormItem label="Name" path="name">
        <NInput v-model:value="modalForm.name" />
      </NFormItem>

      <NFormItem label="Serial No." path="serialNumber">
        <NInput v-model:value="modalForm.serialNumber" />
      </NFormItem>

      <NFormItem label="Category" path="categoryId">
        <NSelect
          v-model:value="modalForm.categoryId"
          :options="categoryOptions"
          :loading="categoriesLoading"
          clearable
          placeholder="Please select the category"
        />
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
