<script setup>
import { computed } from 'vue'

const props = defineProps({
  visible: {
    type: Boolean,
    required: true,
  },
  width: {
    type: String,
    default: '600px',
  },
  title: {
    type: String,
    default: '',
  },
  maskClosable: {
    type: Boolean,
    default: false,
  },
})

const emit = defineEmits(['update:visible'])

const show = computed({
  get: () => props.visible,
  set: (v) => emit('update:visible', v),
})
</script>

<template>
  <n-modal
    v-model:show="show"
    preset="card"
    :style="{ width }"
    :title="title"
    :mask-closable="maskClosable"
  >
    <!-- body -->
    <slot />

    <!-- footer -->
    <template #footer>
      <slot name="footer" />
    </template>
  </n-modal>
</template>
