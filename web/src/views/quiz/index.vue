<script setup lang="ts">
import { computed } from 'vue'
import { NButton } from 'naive-ui'
// hooks
import { useModalStore } from '@/store/modal'
import { useQuizzesQuery } from '@/data/quiz'
// components
import TheIcon from '@/components/icon/TheIcon.vue'
import QuizList from '@/components/quiz/QuizList.vue'
import QuizModal from '@/components/quiz/QuizModal.vue'
import CommonPage from '@/components/page/CommonPage.vue'

// query
const { quizzes, loading } = useQuizzesQuery({})
// store hooks
const modal = useModalStore()
const tableData = computed(() => quizzes.value)

function openCreateModal() {
  modal.open(QuizModal, {
    title: 'Create Quiz',
    props: {
      isEditMode: false,
    },
  })
}
</script>

<template>
  <CommonPage show-footer title="Quiz List">
    <template #action>
      <div>
        <NButton class="float-right mr-15" type="primary" @click="openCreateModal">
          <TheIcon icon="material-symbols:add" :size="18" class="mr-5" />Create new quiz
        </NButton>
      </div>
    </template>
    <QuizList :loading="loading" :table-data="tableData" />
  </CommonPage>
</template>
