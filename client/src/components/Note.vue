<script setup>
import { ref, computed, onMounted } from 'vue'
import { marked } from 'marked'

const props = defineProps({
  id: String,
  title: String,
  body: String,
  index: Number,
})

const bodyRef = ref(null)

const emit = defineEmits(['close', 'loadNote', 'moveLeft', 'moveRight', 'reload'])

const parsedBody = computed(() => {
  return marked.parse(props.body)
})

onMounted(() => {
  const handleJoplinInternalLinks = (e) => {
    // an internal joplin link
    if (e.target.matches('a[href^=":/"]')) {
      e.preventDefault()
      emit('loadNote', e.target.href.replace(/^.*:\//, ''))
      return
    }

    if (e.target.matches('a[href]')) {
      e.preventDefault()
      window.open(e.target.href, '_blank')
      return
    }
  }

  bodyRef.value.addEventListener('click', handleJoplinInternalLinks)
})

</script>

<template>
  <div class="note" draggable>
    <div class="buttons">
      <button @click="emit('close', index)">Close</button>
    </div>
    <h1 class="title">{{ title }}</h1>
    <div ref="bodyRef" class="body" v-html="parsedBody"></div>
  </div>
</template>

<style lang="scss" scoped>
.note {
  display: flex;
  position: relative;
  flex-direction: column;
  flex: none;

  padding: 2rem;

  overflow: auto;
  border: solid #333 1px;

  background-color: #ddd;

  &:hover {
    background-color: #ffe;
  }
}

.buttons {
  position: absolute;
  right: 2.25rem;
  top: 3.25rem;
  display: flex;
}

.title {
  margin-right: 2rem;
}

.body {
  overflow: auto;
}
</style>
