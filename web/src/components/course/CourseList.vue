<script setup lang="ts">
import { h } from 'vue'
import { NButton, NPopconfirm } from 'naive-ui'
// hooks
import { useDeleteCourseMutation } from '@/data/courses'
// utils
import { renderIcon } from '@/utils'
// types
import type { Course } from '@/types'
// hooks
import { useModalStore } from '@/store/modal'
// components
import CourseModal from './CourseModal.vue'

// Define props
const props = defineProps<{
  loading: boolean
  tableData: Course[]
}>()
const modal = useModalStore()

// mutation
const { mutateAsync: deleteCourse } = useDeleteCourseMutation()

function onEdit(course: any) {
  modal.open(CourseModal, {
    title: 'Edit Quiz',
    props: {
      course,
    },
  })
}

async function deleteRow(row: any) {
  await deleteCourse({ id: row.ID })
}

const columns = [
  {
    title: 'Name',
    key: 'name',
    width: 200,
  },
  {
    title: 'Slug',
    key: 'slug',
    ellipsis: true,
    width: 300,
  },
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
