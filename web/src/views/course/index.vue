<script setup lang="ts">
import { computed } from 'vue'
import { NButton } from 'naive-ui'
// hooks
import { useModalStore } from '@/store/modal'
import { useCoursesQuery } from '@/data/courses'
// components
import TheIcon from '@/components/icon/TheIcon.vue'
import CommonPage from '@/components/page/CommonPage.vue'
import CourseList from '@/components/course/CourseList.vue'
import CourseModal from '@/components/course/CourseModal.vue'

// query
const { courses, loading } = useCoursesQuery({})
// store hooks
const modal = useModalStore()
const tableData = computed(() => courses.value)

function openCreateModal() {
  modal.open(CourseModal, {
    title: 'Create Quiz',
  })
}
</script>
<template>
  <CommonPage show-footer title="Course List">
    <template #action>
      <div>
        <NButton class="float-right mr-15" type="primary" @click="openCreateModal">
          <TheIcon icon="material-symbols:add" :size="18" class="mr-5" />Create new course
        </NButton>
      </div>
    </template>
    <CourseList :loading="loading" :table-data="tableData" />
  </CommonPage>
</template>
