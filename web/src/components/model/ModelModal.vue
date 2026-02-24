<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { NForm, NFormItem, NInput, NButton } from 'naive-ui'
// hooks
import { useModalStore } from '@/store/modal'
import { useCategoriesQuery } from '@/data/category'
import { useManufacturersQuery } from '@/data/manufacturer'
import { useCreateModelMutation, useUpdateModelMutation } from '@/data/model'
// types
import type { Model } from '@/types'

const props = defineProps<{
  model?: Model | null
}>()

const modal = useModalStore()

// query
const { categories, loading: categoriesLoading } = useCategoriesQuery({})
const { manufacturers, loading: manufacturersLoading } = useManufacturersQuery({})
// mutations
const { mutateAsync: createModel, isPending: creating } = useCreateModelMutation()
const { mutateAsync: updateModel, isPending: updating } = useUpdateModelMutation()

const modalFormRef = ref()
const modalForm = ref({
  name: '',
  categoryId: '',
  manufacturerId: '',
  modelNumber: '',
})

const categoryOptions = computed(
  () =>
    categories.value?.map((category) => ({
      label: category.name,
      value: category.id,
    })) ?? [],
)

const manufacturerOptions = computed(
  () =>
    manufacturers.value?.map((manufacturer) => ({
      label: manufacturer.name,
      value: manufacturer.id,
    })) ?? [],
)

watch(
  () => props.model,
  (model) => {
    if (!model) {
      // create mode
      modalForm.value = {
        name: '',
        categoryId: '',
        manufacturerId: '',
        modelNumber: '',
      }
      return
    }

    // edit mode
    modalForm.value = {
      name: model.name,
      categoryId: model.categoryId,
      manufacturerId: model.manufacturerId,
      modelNumber: model.modelNumber,
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
    if (props.model) {
      await updateModel({
        id: props.model.id,
        name: modalForm.value.name,
        categoryId: modalForm.value.categoryId,
        manufacturerId: modalForm.value.manufacturerId,
        modelNumber: modalForm.value.modelNumber,
      })
    } else {
      await createModel({
        name: modalForm.value.name,
        categoryId: modalForm.value.categoryId,
        manufacturerId: modalForm.value.manufacturerId,
        modelNumber: modalForm.value.modelNumber,
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

      <NFormItem label="Category" path="categoryId">
        <NSelect
          v-model:value="modalForm.categoryId"
          :options="categoryOptions"
          :loading="categoriesLoading"
          clearable
          placeholder="Please select the category"
        />
      </NFormItem>

      <NFormItem label="Manufacturer" path="manufacturerId">
        <NSelect
          v-model:value="modalForm.manufacturerId"
          :options="manufacturerOptions"
          :loading="manufacturersLoading"
          clearable
          placeholder="Please select the manufacturer"
        />
      </NFormItem>

      <NFormItem label="Model No." path="modelNumber">
        <NInput v-model:value="modalForm.modelNumber" />
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
