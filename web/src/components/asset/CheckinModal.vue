<script setup lang="ts">
import { ref } from 'vue'
import { NForm, NFormItem, NInput, NButton, NDatePicker, type FormRules, NSelect } from 'naive-ui'
// hooks
import { useModalStore } from '@/store/modal'
import { useCheckinAssetMutation } from '@/data/asset'
// types
import type { Asset } from '@/types'
import { AssetStatus } from '@/types/enum'

const props = defineProps<{
  asset: Asset
}>()

const modal = useModalStore()

// mutations
const { mutateAsync: checkinAsset, isPending } = useCheckinAssetMutation()

const modalFormRef = ref()
const modalForm = ref({
  assetName: '',
  userId: null,
  checkinDate: null,
  status: '',
  notes: '',
})

const validationRules: FormRules = {
  userId: [{ required: true, type: 'number', message: 'User is required', trigger: ['blur'] }],
  checkinDate: [
    { required: true, type: 'number', message: 'Checkin Date is required', trigger: ['blur'] },
  ],
}

const statusOptions = [
  {
    label: 'Pending',
    value: AssetStatus.AssetPending,
  },
  {
    label: 'Ready to Deploy',
    value: AssetStatus.AssetReadyToDeploy,
  },
  {
    label: 'Archived',
    value: AssetStatus.AssetArchived,
  },
  {
    label: 'Broken - Not Fixable',
    value: AssetStatus.AssetBroken,
  },
  {
    label: 'Lost/Stolen',
    value: AssetStatus.AssetLostStolen,
  },
]

function toISO(value: number | null) {
  return value ? new Date(value).toISOString() : null
}

async function handleCheckout() {
  modalFormRef.value?.validate(async (errors: Error) => {
    if (errors) return

    console.log('modalForm:', modalForm.value.status)

    await checkinAsset({
      assetName: modalForm.value.assetName,
      assetId: props.asset.id,
      checkinDate: toISO(modalForm.value.checkinDate),
      status: modalForm.value.status,
      notes: modalForm.value.notes,
    })

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
      <NFormItem label="Asset Name" path="assetName">
        <NInput v-model:value="modalForm.assetName" />
      </NFormItem>

      <NFormItem label="Checkin Date" path="checkinDate">
        <NDatePicker v-model:value="modalForm.checkinDate" type="datetime" clearable />
      </NFormItem>

      <NFormItem label="Status" path="status">
        <NSelect
          v-model:value="modalForm.status"
          :options="statusOptions"
          clearable
          placeholder="Please select the status"
        />
      </NFormItem>

      <NFormItem label="Notes" path="note">
        <NInput v-model:value="modalForm.notes" />
      </NFormItem>
    </NForm>

    <div flex justify-end>
      <NButton @click="modal.close">Cancel</NButton>
      <NButton type="primary" class="ml-16" :loading="isPending" @click="handleCheckout">
        Save
      </NButton>
    </div>
  </div>
</template>
