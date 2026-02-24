<script setup lang="ts">
import { computed, ref, watch } from 'vue'
// hooks
import { useModelsQuery } from '@/data/model'
import { useModalStore } from '@/store/modal'
import { useCreateAssetMutation, useUpdateAssetMutation } from '@/data/asset'
// types
import { AssetStatus } from '@/types/enum'
import type { Asset } from '@/types'
// components
import ModelModal from '../model/ModelModal.vue'

const props = defineProps<{
  asset?: Asset | null
}>()

const modal = useModalStore()
// query
const { models, loading: modelsLoading } = useModelsQuery({})
// mutations
const { mutateAsync: createAsset, isPending: creatingAsset } = useCreateAssetMutation()
const { mutateAsync: updateAsset, isPending: updatingAsset } = useUpdateAssetMutation()

const rules = {
  name: [{ required: true, message: 'Name is required', trigger: ['blur'] }],
}
const modalFormRef = ref()
const modalForm = ref({
  name: '',
  tag: '',
  serialNumber: '',
  model: '',
  status: '',
  description: '',
})

watch(
  () => props.asset,
  (asset) => {
    if (!asset) {
      // create mode
      // modalForm.value = {
      //   name: '',
      //   serialNumber: '',
      // }
      return
    }

    // edit mode
    modalForm.value = {
      name: asset.name,
      tag: asset.tag,
      serialNumber: asset.serialNumber,
      status: asset.status,
      description: asset.description,
      model: asset?.model?.id,
    }
  },
  { immediate: true },
)

const modelOptions = computed(
  () =>
    models.value?.map((model) => ({
      label: model.name,
      value: model.id,
    })) ?? [],
)

// const generalOptions = ['groode', 'veli good', 'emazing', 'lidiculous'].map((v) => ({
//   label: v,
//   value: v,
// }))

const statusOptions = [
  {
    label: 'Pending',
    value: AssetStatus.AssetPending,
  },
  {
    label: 'Ready To Deploy',
    value: AssetStatus.AssetReadyToDeploy,
  },
  {
    label: 'Archived',
    value: AssetStatus.AssetArchived,
  },
  {
    label: 'Broken',
    value: AssetStatus.AssetBroken,
  },
  {
    label: 'Lost',
    value: AssetStatus.AssetLostStolen,
  },
]

function handleAddNewModel() {
  modal.open(ModelModal, {
    title: 'Add Model',
  })
}

async function handleSave() {
  modalFormRef.value?.validate(async (errors: Error) => {
    if (errors) return
    if (props.asset) {
      await updateAsset({
        id: props.asset.id,
        name: modalForm.value.name,
        tag: modalForm.value.tag,
        serialNumber: modalForm.value.serialNumber,
        modelId: modalForm.value.model,
        status: modalForm.value.status,
        description: modalForm.value.description,
      })
    } else {
      await createAsset({
        name: modalForm.value.name,
        tag: modalForm.value.tag,
        serialNumber: modalForm.value.serialNumber,
        modelId: modalForm.value.model,
        status: modalForm.value.status,
        description: modalForm.value.description,
      })
    }
    modal.close()
  })
}
</script>

<template>
  <n-form
    ref="modalFormRef"
    :model="modalForm"
    :rules="rules"
    label-placement="left"
    require-mark-placement="right-hanging"
    label-width="auto"
    :style="{ maxWidth: '720px' }"
    @submit.prevent="handleSave"
  >
    <n-form-item label="Name" path="name">
      <n-input v-model:value="modalForm.name" placeholder="Input" />
    </n-form-item>
    <n-form-item label="Asset Tag" path="tag">
      <n-input v-model:value="modalForm.tag" placeholder="Input" />
    </n-form-item>
    <n-form-item label="Serial" path="serialNumber">
      <n-input v-model:value="modalForm.serialNumber" placeholder="Input" />
    </n-form-item>
    <n-form-item label="Model" path="model">
      <n-select v-model:value="modalForm.model" placeholder="Select" :options="modelOptions" />
      <n-button style="margin-left: 12px" @click="handleAddNewModel()"> New </n-button>
    </n-form-item>
    <n-form-item label="Status" path="status">
      <n-select v-model:value="modalForm.status" placeholder="Select" :options="statusOptions" />
    </n-form-item>
    <n-form-item :span="12" label="Description" path="textareaValue">
      <n-input
        v-model:value="modalForm.description"
        placeholder="Textarea"
        type="textarea"
        :autosize="{
          minRows: 3,
          maxRows: 5,
        }"
      />
    </n-form-item>
    <n-form-item>
      <n-button
        type="primary"
        attr-type="submit"
        :loading="creatingAsset || updatingAsset"
        :disabled="creatingAsset || updatingAsset"
      >
        Save Asset
      </n-button>
    </n-form-item>
  </n-form>
</template>
