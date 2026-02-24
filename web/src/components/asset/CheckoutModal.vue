<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { NForm, NFormItem, NInput, NButton, NDatePicker, type FormRules } from 'naive-ui'
// hooks
import { useUsersQuery } from '@/data/user'
import { useModalStore } from '@/store/modal'
import { useCheckoutAssetMutation } from '@/data/asset'
// types
import type { Asset, User } from '@/types'

const props = defineProps<{
  asset: Asset
}>()

const modal = useModalStore()

// query
const { users, loading: usersLoading } = useUsersQuery({})

// mutations
const { mutateAsync: checkoutAsset, isPending } = useCheckoutAssetMutation()

const modalFormRef = ref()
const modalForm = ref({
  assetName: '',
  userId: null,
  checkoutDate: null,
  expectedCheckinDate: null,
  notes: '',
})

// watch(
//   () => props.asset,
//   (asset) => {
//     if (!asset) {
//       // create mode
//       modalForm.value = {
//         name: '',
//         serialNumber: '',
//         categoryId: '',
//         note: '',
//       }
//       return
//     }

//     // edit mode
//     modalForm.value = {
//       name: asset.name,
//       serialNumber: asset.serialNumber,
//       categoryId: asset.categoryId,
//     }
//   },
//   { immediate: true },
// )

const userOptions = computed(
  () =>
    users.value?.map((user: User) => ({
      label: user.email,
      value: user.id,
    })) ?? [],
)

const validationRules: FormRules = {
  userId: [{ required: true, type: 'number', message: 'User is required', trigger: ['blur'] }],
  checkoutDate: [
    { required: true, type: 'number', message: 'Checkout Date is required', trigger: ['blur'] },
  ],
}

function toISO(value: number | null) {
  return value ? new Date(value).toISOString() : null
}

async function handleCheckout() {
  modalFormRef.value?.validate(async (errors: Error) => {
    if (errors) return

    await checkoutAsset({
      assetName: modalForm.value.assetName,
      assetId: props.asset.id,
      userId: modalForm.value.userId,
      checkoutDate: toISO(modalForm.value.checkoutDate),
      expectedCheckinDate: toISO(modalForm.value.expectedCheckinDate),
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

      <NFormItem label="User" path="userId">
        <NSelect
          v-model:value="modalForm.userId"
          :options="userOptions"
          :loading="usersLoading"
          clearable
          placeholder="Please select the category"
        />
      </NFormItem>

      <NFormItem label="Checkout Date" path="checkoutDate">
        <NDatePicker v-model:value="modalForm.checkoutDate" type="datetime" clearable />
      </NFormItem>

      <NFormItem label="Expected Check-in Date" path="expectedCheckinDate">
        <NDatePicker v-model:value="modalForm.expectedCheckinDate" type="datetime" clearable />
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
