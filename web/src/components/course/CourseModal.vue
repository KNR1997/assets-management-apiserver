<script setup lang="ts">
import slugify from 'slugify'
import { ref, watch } from 'vue'
import { NForm, NFormItem, NInput, NButton } from 'naive-ui'
// hooks
import { useModalStore } from '@/store/modal'
import { useCreateCourseMutation, useUpdateCourseMutation } from '@/data/courses'
import type { Course } from '@/types'

const props = defineProps<{
  course?: Course | null
}>()

const modal = useModalStore()

// mutations
const { mutateAsync: createCourse, isPending: creating } = useCreateCourseMutation()
const { mutateAsync: updateCourse, isPending: updating } = useUpdateCourseMutation()

const modalFormRef = ref()
const modalForm = ref({
  name: '',
  slug: '',
})

watch(
  () => props.course,
  (course) => {
    if (!course) {
      // create mode
      modalForm.value = {
        name: '',
        slug: '',
      }
      return
    }

    // edit mode
    modalForm.value = {
      name: course.name,
      slug: course.slug,
    }
  },
  { immediate: true },
)

watch(
  () => modalForm.value.name,
  (name) => {
    modalForm.value.slug = slugify(name ?? '', {
      lower: true,
      strict: true,
      trim: true,
    })
  },
)

const validateAddCourse = {
  name: [{ required: true, message: 'Name is required', trigger: ['blur'] }],
  slug: [
    { required: true, message: 'Slug is required', trigger: ['blur'] },
    {
      pattern: /^[a-z0-9]+(?:-[a-z0-9]+)*$/,
      message: 'Slug must be URL-safe',
      trigger: ['input'],
    },
  ],
}

async function handleSave() {
  modalFormRef.value?.validate(async (errors: Error) => {
    if (errors) return
    if (props.course) {
      await updateCourse({
        id: props.course.ID,
        name: modalForm.value.name,
        slug: modalForm.value.slug,
      })
    } else {
      await createCourse({
        name: modalForm.value.name,
        slug: modalForm.value.slug,
      })
    }
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
      :rules="validateAddCourse"
    >
      <NFormItem label="Name" path="name">
        <NInput v-model:value="modalForm.name" />
      </NFormItem>

      <NFormItem label="Slug" path="slug">
        <NInput v-model:value="modalForm.slug" />
      </NFormItem>
    </NForm>

    <div flex justify-end>
      <NButton @click="modal.close">Cancel</NButton>
      <NButton type="primary" class="ml-16" :loading="creating || updating" @click="handleSave">
        Save
      </NButton>
    </div>
  </div>
</template>
