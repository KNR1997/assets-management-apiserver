import { defineStore } from 'pinia'
import { activeTag, tags, WITHOUT_TAG_PATHS } from './helpers'
import { router } from '@/router'
import { lStorage } from '@/utils'

export const useTagsStore = defineStore('tag', {
  state() {
    return {
      tags: tags || [],
      activeTag: activeTag || '',
    }
  },
  getters: {
    // @ts-ignore
    activeIndex() {
      // @ts-ignore
      return this.tags.findIndex((item) => item.path === this.activeTag)
    },
  },
  actions: {
    setActiveTag(path: any) {
      this.activeTag = path
      // @ts-ignore
      lStorage.set('activeTag', path)
    },
    setTags(tags: any) {
      this.tags = tags
      // @ts-ignore
      lStorage.set('tags', tags)
    },
    addTag(tag = {}) {
      // @ts-ignore
      this.setActiveTag(tag.path)
      // @ts-ignore
      if (WITHOUT_TAG_PATHS.includes(tag.path) || this.tags.some((item: any) => item.path === tag.path))
        return
      this.setTags([...this.tags, tag])
    },
    removeTag(path: any) {
      if (path === this.activeTag) {
        if (this.activeIndex > 0) {
          router.push(this.tags[this.activeIndex - 1].path)
        } else {
          router.push(this.tags[this.activeIndex + 1].path)
        }
      }
      this.setTags(this.tags.filter((tag: any) => tag.path !== path))
    },
    // @ts-ignore
    removeOther(curPath = this.activeTag) {
      this.setTags(this.tags.filter((tag: any) => tag.path === curPath))
      if (curPath !== this.activeTag) {
        router.push(this.tags[this.tags.length - 1].path)
      }
    },
    removeLeft(curPath: any) {
      const curIndex = this.tags.findIndex((item: any) => item.path === curPath)
        // @ts-ignore
      const filterTags = this.tags.filter((item: any, index: any) => index >= curIndex)
      this.setTags(filterTags)
      if (!filterTags.find((item: any) => item.path === this.activeTag)) {
        router.push(filterTags[filterTags.length - 1].path)
      }
    },
    removeRight(curPath: any) {
      const curIndex = this.tags.findIndex((item: any) => item.path === curPath)
        // @ts-ignore
      const filterTags = this.tags.filter((item: any, index: any) => index <= curIndex)
      this.setTags(filterTags)
      if (!filterTags.find((item: any) => item.path === this.activeTag)) {
        router.push(filterTags[filterTags.length - 1].path)
      }
    },
    resetTags() {
      this.setTags([])
      this.setActiveTag('')
    },
  },
})
