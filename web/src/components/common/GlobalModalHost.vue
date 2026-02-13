<script setup lang="ts">
import { useModalStore } from '@/store/modal'
import Modal from '@/components/common/Modal.vue'

const modal = useModalStore()
</script>

<template>
  <Modal v-model:visible="modal.visible" :title="modal.title">
    <component :is="modal.component" v-bind="modal.props" />
  </Modal>
</template>
