<template>
  <div flex items-center>
    <MenuCollapse />
    <BreadCrumb ml-15 hidden sm:block />
  </div>
  <div ml-auto flex items-center>
    <Languages />
    <ThemeMode />
    <GithubSite />
    <FullScreen />
    <UserAvatar />
  </div>
</template>

<script setup>
import MenuCollapse from "./components/MenuCollapse.vue";
import BreadCrumb from "./components/BreadCrumb.vue";
import Languages from "./components/Languages.vue";
import ThemeMode from "./components/ThemeMode.vue";
import GithubSite from "./components/GithubSite.vue";
import FullScreen from "./components/FullScreen.vue";
import UserAvatar from "./components/UserAvatar.vue";
</script>
