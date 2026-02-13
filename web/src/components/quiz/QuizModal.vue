<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import type { FormRules } from 'naive-ui'
import { NForm, NFormItem, NInput, NSelect, NDatePicker } from 'naive-ui'
// types
import type { Quiz } from '@/types'
// hooks
import { useModalStore } from '@/store/modal'
import { useCoursesQuery } from '@/data/courses'
import { useCreateQuizMutation, useUpdateQuizMutation } from '@/data/quiz'

const props = defineProps<{
  quiz?: Quiz | null
}>()

const modal = useModalStore()

// query
const { courses, loading: coursesLoading } = useCoursesQuery({})

// mutation
const { mutateAsync: createQuiz, isPending: creating } = useCreateQuizMutation()
const { mutateAsync: updateQuiz, isPending: updating } = useUpdateQuizMutation()

interface QuizForm {
  courseId: number | null
  name: string
  date: number | null
}

const modalFormRef = ref()
const modalForm = ref<QuizForm>({
  courseId: null,
  name: '',
  date: null,
})

watch(
  () => props.quiz,
  (quiz) => {
    if (!quiz) {
      // create mode
      // modalForm.value = {
      //   name: '',
      // }
      return
    }

    // edit mode
    modalForm.value = {
      courseId: quiz.courseId,
      name: quiz.name,
      date: quiz.date,
    }
  },
  { immediate: true },
)

const courseOptions = computed(
  () =>
    courses.value?.map((course) => ({
      label: course.name,
      value: course.ID,
    })) ?? [],
)

const validationRules: FormRules = {
  courseId: [
    {
      required: true,
      type: 'number',
      message: 'Course is required',
      trigger: 'change',
    },
  ],
  name: [
    {
      required: true,
      message: 'Name is required',
      trigger: ['input', 'blur'],
    },
  ],
  date: [
    {
      required: true,
      type: 'number',
      message: 'Quiz date is required',
      trigger: 'change',
    },
  ],
}

async function handleSave() {
  await modalFormRef.value?.validate()
  if (!modalForm.value.courseId) return
  if (props.quiz) {
    await updateQuiz({
      // @ts-ignore
      id: props.quiz.ID, // todo -> fix
      courseId: props.quiz.courseId,
      name: modalForm.value.name,
      dateTime: new Date(modalForm.value.date!).toISOString(),
    })
  } else {
    await createQuiz({
      courseId: modalForm.value.courseId,
      name: modalForm.value.name,
      dateTime: new Date(modalForm.value.date!).toISOString(),
    })
  }
  modal.close()
}
</script>

<template>
  <NForm
    ref="modalFormRef"
    label-placement="left"
    label-align="left"
    :label-width="80"
    :model="modalForm"
    :rules="validationRules"
  >
    <NFormItem label="Course" path="courseId">
      <NSelect
        v-model:value="modalForm.courseId"
        :options="courseOptions"
        :loading="coursesLoading"
        clearable
        placeholder="Please select the course"
      />
    </NFormItem>
    <NFormItem label="Name" path="name">
      <NInput v-model:value="modalForm.name" clearable placeholder="Please enter the name" />
    </NFormItem>
    <NFormItem label="Date" path="date">
      <NDatePicker
        v-model:value="modalForm.date"
        type="date"
        clearable
        placeholder="Select quiz date"
      />
    </NFormItem>
  </NForm>
  <div flex justify-end>
    <NButton @click="modal.close">Cancel</NButton>
    <NButton type="primary" class="ml-16" :loading="creating || updating" @click="handleSave">
      Save
    </NButton>
  </div>
</template>
