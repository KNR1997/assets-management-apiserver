<template>
  <AppPage :show-footer="true">
    <div flex-1>
      <n-card rounded-10>
        <div flex items-center justify-between>
          <div flex items-center>
            <img rounded-full width="60" src="/tauri.svg" />
            <div ml-10>
              <p text-20 font-semibold>hello, admin</p>
              <p mt-5 text-14 op-60>Today is another day full of energy!</p>
            </div>
          </div>
          <!-- <n-space :size="12" :wrap="false">
            <n-statistic
              v-for="item in statisticData"
              :key="item.id"
              v-bind="item"
            ></n-statistic>
          </n-space> -->
        </div>
      </n-card>
      <!-- 
      <n-card size="small" :segmented="true" mt-15 rounded-10>
        <template #header-extra>
          <n-button text type="primary">More</n-button>
        </template>
        <div flex flex-wrap justify-between>
          <n-card
            v-for="course in courses"
            :key="course.id"
            class="mb-10 mt-10 w-300 cursor-pointer"
            hover:card-shadow
            :title="course.code + ' ' + course.name"
            size="small"
            @click="openCourseModal(course)"
          >
            <p op-60>{{ dummyText }}</p>
            <p mt-15>Quizzes: 1/5</p>
          </n-card>
        </div>
      </n-card> -->
      <!-- <n-modal v-model:show="showModal" preset="card" style="width: 600px" title="Course Details">
        <div v-if="selectedCourse">
          <p><b>Code:</b> {{ selectedCourse.code }}</p>
          <p><b>Name:</b> {{ selectedCourse.name }}</p>
          <p><b>Description:</b> {{ selectedCourse.description || '—' }}</p>
        </div>
      </n-modal> -->
    </div>
  </AppPage>
</template>

<script setup>
import { ref } from 'vue'
import AppPage from '@/components/page/AppPage.vue'
import { useCoursesQuery } from '~/src/data/courses'

// const { courses } = useCoursesQuery({})

console.log('health checking-------------')

fetch('api/health')

// import { useCoursesQuery } from '~/src/data/courses'

// const { courses } = useCoursesQuery({})

// const showModal = ref(false)
// const selectedCourse = ref(null)

// const openCourseModal = (course) => {
//   selectedCourse.value = course
//   showModal.value = true
// }

// async function getCourseList(params) {
//   const res = await api.getCourseList()
//   console.log('res------------: ', res)
// }

// getCourseList()

const dummyText = 'A lightweight admin management template based on Vue 3.0, FastAPI, and Naive UI.'
</script>
