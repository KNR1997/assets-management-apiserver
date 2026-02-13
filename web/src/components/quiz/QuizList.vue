<script setup lang="ts">
import { h } from 'vue'
import { NButton, NPopconfirm } from 'naive-ui'
// hooks
import { useDeleteQuizMutation } from '@/data/quiz'
// utils
import { renderIcon } from '@/utils'
import type { Quiz } from '@/types'
import QuizModal from './QuizModal.vue'
import { useModalStore } from '@/store/modal'

// Define props
const props = defineProps<{
  loading: boolean
  tableData: Quiz[]
}>()
// store hooks
const modal = useModalStore()
// mutation
const { mutateAsync: deleteQuiz } = useDeleteQuizMutation()

function onEdit(quiz: any) {
  modal.open(QuizModal, {
    title: 'Edit Quiz',
    props: {
      quiz,
    },
  })
}

async function deleteRow(row: any) {
  await deleteQuiz({ id: row.ID })
}

const columns = [
  {
    title: 'Name',
    key: 'name',
    width: 200,
  },
  {
    title: 'Course',
    key: 'course.name',
    ellipsis: true,
    width: 300,
  },
  // {
  //   title: 'Quiz Date',
  //   key: 'date',
  //   ellipsis: true,
  //   width: 300,
  // },
  {
    title: 'Actions',
    key: 'actions',
    width: 160,
    render(row: any) {
      return [
        h(
          NButton,
          {
            size: 'small',
            type: 'primary',
            class: 'mr-5',
            onClick: () => onEdit(row),
          },
          { default: () => 'Edit', icon: renderIcon('material-symbols:edit', { size: 16 }) },
        ),
        h(
          NPopconfirm,
          {
            onPositiveClick: () => deleteRow(row),
          },
          {
            trigger: () =>
              h(
                NButton,
                {
                  size: 'small',
                  type: 'error',
                },
                {
                  default: () => 'Delete',
                  icon: renderIcon('material-symbols:delete-outline', { size: 16 }),
                },
              ),
            default: () => 'Are you sure?',
          },
        ),
      ]
    },
  },
]
</script>

<template>
  <n-data-table :loading="loading" :columns="columns" :data="tableData" />
</template>
