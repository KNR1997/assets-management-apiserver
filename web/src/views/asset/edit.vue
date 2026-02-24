<script setup lang="ts">
import { computed } from 'vue'
// hooks
import { useRoute } from 'vue-router'
import { useAssetQuery } from '@/data/asset'
// components
import CommonPage from '@/components/page/CommonPage.vue'
import AssetForm from '@/components/asset/AssetForm.vue'

const route = useRoute()
const assetId = computed(() => route.params.id as string)

const { asset, loading, error } = useAssetQuery({ id: assetId.value })

</script>

<template>
  <CommonPage show-footer title="Edit Asset">
    <AssetForm :asset="asset" />
  </CommonPage>
</template>
