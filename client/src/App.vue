<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { fromUint8Array, toUint8Array } from 'js-base64'
import draggable from 'vuedraggable'

import Note from './components/Note.vue'
import NoteSearch from './components/NoteSearch.vue'

const results = ref([])
const notes = ref([])
const searchFocus = ref(+new Date)

const drag = ref(false)

const router = useRouter()
const route = useRoute()

function base64ToUUID(base64) {
  return toUint8Array(base64).reduce((a, n) => (
    a + ('0' + n.toString(16)).substr(-2)
  ), "")
}

function uuidToBase64(uuid) {
  const bytes = []
  for (let i = 0; i < uuid.length; i += 2) {
    bytes.push(parseInt(uuid.substr(i, 2), 16))
  }

  return fromUint8Array(Uint8Array.from(bytes), true)
}

if (route.query.ids) {
  const ids = route.query.ids.split(',')
  router.push({query: {ids: ''}})

  const loadNotes = async () => {
    for (let i = 0; i < ids.length; i++) {
      await loadNote(base64ToUUID(ids[i]))
    }
  }

  loadNotes()
}

watch(notes, (newNotes) => {
  newNotes.forEach(n => {
    uuidToBase64(n.id)
  })

  router.push({
    query: { ids: newNotes.map(n => uuidToBase64(n.id)).join(',') }
  })

  window.document.title = newNotes.map(n => n.title).join(', ')
})

async function performSearch(query) {
  if (query.length < 2) {
    results.value = []
    return
  }

  const response = await fetch(`/api/v1/search?q=${encodeURIComponent(query)}`)
  const responseJSON = await response.json()
  results.value = responseJSON
}

async function loadNoteData(id) {
  const response = await fetch(`/api/v1/note/${id}`)
  const responseJSON = await response.json()
  return responseJSON
}

async function loadNote(id) {
  try {
    const noteData = await loadNoteData(id)

    if (noteData.id === "") {
      console.error(`no note data for note ${id}`)
      return
    }

    notes.value = notes.value.concat([noteData])

    results.value = []
    focusSearch()
  } catch (e) {
    console.error(`unable to load note ${id}`)
  }
}

function moveLeft(index) {
  const newNotes = []
  for (let i = 0; i < notes.value.length; i++) {
    if (i === index) {
      newNotes[i] = notes.value[i - 1]
      newNotes[i - 1] = notes.value[i]
    } else {
      newNotes[i] = notes.value[i]
    }
  }

  notes.value = newNotes
}

function moveRight(index) {
  const newNotes = []
  for (let i = 0; i < notes.value.length; i++) {
    if (i === index + 1) {
      newNotes[i] = notes.value[i - 1]
      newNotes[i - 1] = notes.value[i]
    } else {
      newNotes[i] = notes.value[i]
    }
  }

  notes.value = newNotes
}

function closeNote(index) {
  notes.value = notes.value.filter((_, idx) => idx !== index)
}

function focusSearch() {
  searchFocus.value++
}

async function reloadNote(index) {
  const id = notes.value[index].id
  const newNote = await loadNoteData(id)

  notes.value = notes.value.map(n => {
    if (n.id === id) {
      return newNote
    } else {
      return n
    }
  })
}

onMounted(() => {
  document.addEventListener('keydown', (e) => {
    if (e.target.matches('input')) {
      return
    }

    if (e.keyCode == 83) { // s
      e.preventDefault()
      focusSearch()
    }
  })

  setInterval(async () => {
    for (let i = 0; i < notes.value.length; ++i) {
      await reloadNote(i)
    }
  }, 5000)
})
</script>

<template>
  <draggable
    v-model="notes"
    item-key="id"
    id="notes-container"
  >
    <template #item="{element, index}">
      <Note
        v-bind="element"
        :index="index"

        @loadNote="loadNote"
        @close="closeNote"
      />
    </template>
    <template #footer>
      <NoteSearch
        @search="performSearch"
        @select="loadNote"
        :searchFocus="searchFocus"
        ref="noteSearch"
        :results="results"
      />
    </template>
  </draggable>
</template>

<style lang="scss">
html, body, #app {
  height: 100vh;
  margin: 0;
  padding: 0;
}

html {
  box-sizing: border-box;
}

#notes-container {
  display: flex;

  height: 100vh;
  overflow-x: auto;
  overflow-y: none;

  > div {
    width: calc(35rem - 4rem);
    padding: 2rem;
    height: calc(100vh - 6rem);
  }
}
</style>
